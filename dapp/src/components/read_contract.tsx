import React, {useEffect, useState } from 'react';
import {Text} from '@chakra-ui/react'
//import {ERC20ABI as abi} from 'abi/ERC20ABI'
import {abi} from '../../../artifacts/contracts/MetaMarketplace.sol/MetaMarketplace.json'

import {Contract, ethers} from 'ethers'


interface Props {
    addressContract: string,
    currentAccount: string | undefined
}

declare let window: any;

export default function ReadLastOffer(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  var [currency, setCurrency] = useState(0) // TODO: fix it to work as input option
  var [price,setPrice] = useState<string>("")
  var [human_number,setHuman_number] = useState<string>("")
  var [token_id,setTokenId] = useState(0)
  var [collection_address, setCollectionAddress] = useState<string>("")

  // called only once, use it as constructor
  useEffect( () => {
    if(!window.ethereum) return

    const queryParams = new URLSearchParams(location.search);

    var ac_q = queryParams.get('collection_contract_address'); // erc20 to approve
    var ac_a = ethers.utils.getAddress(ac_q);
    setCollectionAddress(ac_a);


    var token_id_q = queryParams.get('token_id'); // amount MUST be in wei format
    let token_id_n : number = +token_id_q;
    setTokenId(token_id_n);

    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const MetaMarketplace = new ethers.Contract(addressContract, abi, provider);
    MetaMarketplace.getLastPrice(collection_address,token_id).then((result:string)=>{
        console.log(result)
        
       // setOwnerValue(result)
       // setPrice(result[0])
       // setCurrency(result[1])
    }).catch('error', console.error)



    //called only once
  },[])  

  return (


    <div>
        <Text><b>Marketplace address</b>: {addressContract}</Text>
        <Text><b>Last price</b>:{price}</Text>
        <Text><b>Last currency</b>:{currency}</Text>
    </div>

  )
}