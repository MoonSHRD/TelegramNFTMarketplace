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


export default function AcceptBuyOffer(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  const marketAddress = props.marketAddress
  //var [user_id, setUserId] = useState(0)
  //var [user_name, setUserName] = useState<string>("")
  var [currency, setCurrency] = useState(0) // TODO: fix it to work as input option
  var [price,setPrice] = useState<string>("")
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

  var p = queryParams.get('price');
  var a_ether = ethers.utils.formatEther(p);
  setHuman_number(a_ether);
  setPrice(p);

  var ac_q = queryParams.get('collection_contract_address'); // erc20 to approve
  var ac_a = ethers.utils.getAddress(ac_q);
  setCollectionAddress(ac_a);
  
  var c = queryParams.get('currency');
  var c_int = parseInt(c);
  setCurrency(c_int);


  /*
  const MetaMarketplace = new ethers.Contract(addressContract, abi, provider);
  MetaMarketplace.getLastPrice(collection_address,token_id_q).then((result:string)=>{
        console.log(result)
        
        setPrice(result)
    }).catch('error', console.error)
*/
 
  



  
  
  
  
  }, []);
  

  async function acceptBuyOffer(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const user_address = signer._address
    const MetaMarketplace:Contract = new ethers.Contract(addressContract, abi, signer)
    console.log("token id to interact raw: ", token_id)
    // var token_id_uint = ethers.utils.
 
    MetaMarketplace.acceptBuyOffer(collection_address,token_id,currency)
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("accept buy offer receipt", receipt)})
        })
         .catch((e:Error) => console.log(e))
     }


  
  //const handleChange = (value:string) => setUserId(value)   
  
  return (
    <form onSubmit={acceptBuyOffer}>
    <FormControl>
      <FormLabel >Sell your nft: </FormLabel>
      <Input id="token_id" type="number" required  onChange={(e) => setTokenId(parseInt(e.target.value))} value={token_id} my={3}/>
     
      
     <Input id="collection_contract_address" type="text" required  onChange={(e) => setCollectionAddress(e.target.value)} value={collection_address} my={3}/>

      <div>
        <Text><b>Token id to sell</b>:{token_id}</Text>
        <Text><b>Price</b>:{price}</Text>
        <Text><b>Currency</b>:{currency}</Text>
    </div>
      <Button type="submit" isDisabled={!currentAccount}>Sell</Button>
    </FormControl>
    </form>
  )
}