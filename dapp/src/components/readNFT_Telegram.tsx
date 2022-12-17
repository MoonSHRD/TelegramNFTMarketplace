import React, {useEffect, useState } from 'react';
import {Text, Image} from '@chakra-ui/react'
//import {ERC20ABI as abi} from 'abi/ERC20ABI'
import {abi} from '../../../artifacts/@openzeppelin/contracts/token/ERC721/extensions/IERC721Metadata.sol/IERC721Metadata.json'

import {Contract, ethers} from 'ethers'


interface Props {
    addressContract: string,
    currentAccount: string | undefined
}

declare let window: any;

export default function ReadLastOffer(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount


  var [token_id,setTokenId] = useState(0)
  var [uri,setUri] = useState<string>("")

  // called only once, use it as constructor
  useEffect( () => {
    if(!window.ethereum) return

    const queryParams = new URLSearchParams(location.search);

    var token_id_q = queryParams.get('token_id'); // amount MUST be in wei format
    let token_id_n : number = +token_id_q;
    setTokenId(token_id_n);

    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const ERCMeta = new ethers.Contract(addressContract, abi, provider);
    ERCMeta.tokenURI(token_id).then((result:string)=>{
        console.log("uri:");
        
        console.log(result);
        setUri(result);
        
       // setOwnerValue(result)
       // setPrice(result[0])
       // setCurrency(result[1])
    }).catch('error', console.error)




    //called only once
  },[])  

  return (


    <div>
        <Text><b>NFT address</b>: {addressContract}</Text>
        <Text><b>Token ID</b>: {token_id}</Text>
        <Image src={uri}></Image>
    </div>

  )
}