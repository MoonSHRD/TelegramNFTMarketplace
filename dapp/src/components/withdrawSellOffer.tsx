import React, { useState, useEffect } from 'react'
import {Button, Input , NumberInput,  NumberInputField,  FormControl,  FormLabel, Text, Select } from '@chakra-ui/react'
import {ethers} from 'ethers'
import {parseEther } from 'ethers/lib/utils'
import {abi} from '../../../artifacts/contracts/MetaMarketplace.sol/MetaMarketplace.json'
import { Contract } from "ethers"
import { TransactionResponse,TransactionReceipt } from "@ethersproject/abstract-provider"


interface Props {
    addressContract: string,
    currentAccount: string | undefined
    marketAddress: string
}

declare let window: any;


export default function WithdrawSellOffer(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  const marketAddress = props.marketAddress
  //var [user_id, setUserId] = useState(0)
  //var [user_name, setUserName] = useState<string>("")


  var [human_number,setHuman_number] = useState<string>("")
  var [token_id,setTokenId] = useState(0)
  var [collection_address, setCollectionAddress] = useState<string>("")


  useEffect(() => {
  const queryParams = new URLSearchParams(location.search);

    if(!window.ethereum) return

  const provider = new ethers.providers.Web3Provider(window.ethereum)
  
  var token_id_q = queryParams.get('token_id'); // amount MUST be in wei format
  let token_id_n : number = +token_id_q;
  setTokenId(token_id_n);


  var ac_q = queryParams.get('collection_contract_address'); // erc20 to approve
  var ac_a = ethers.utils.getAddress(ac_q);
  setCollectionAddress(ac_a);
  



  /*
  const MetaMarketplace = new ethers.Contract(addressContract, abi, provider);
  MetaMarketplace.getLastPrice(collection_address,token_id_q).then((result:string)=>{
        console.log(result)
        
        setPrice(result)
    }).catch('error', console.error)
*/
 
  



  
  
  
  
  }, []);
  

  async function withdrawSellOffer(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const user_address = signer._address
    const MetaMarketplace:Contract = new ethers.Contract(addressContract, abi, signer)
    console.log("token id to interact raw: ", token_id)
    // var token_id_uint = ethers.utils.
 
    MetaMarketplace.withdrawSellOffer(collection_address,token_id)
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("withdraw sell offer receipt", receipt)})
        })
         .catch((e:Error) => console.log(e))
     }


  
  //const handleChange = (value:string) => setUserId(value)   
  
  return (
    <form onSubmit={withdrawSellOffer}>
    <FormControl>
      <FormLabel >Sell your nft: </FormLabel>
      <Input id="token_id" type="number" required  onChange={(e) => setTokenId(parseInt(e.target.value))} value={token_id} my={3}/>
     <Input id="collection_contract_address" type="text" required  onChange={(e) => setCollectionAddress(e.target.value)} value={collection_address} my={3}/>
      <div>
        <Text><b>Token id to sell</b>:{token_id}</Text>
    </div>
      <Button type="submit" isDisabled={!currentAccount}>Put off sale</Button>
    </FormControl>
    </form>
  )
}