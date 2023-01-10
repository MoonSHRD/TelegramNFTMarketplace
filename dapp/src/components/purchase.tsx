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

///buy_telegram?token_id=4&price=1000&currency=5&collection_contract_address=0x2b1e0A2b16AB524Ad3f4273d1ecB63FC3b9cB58C
export default function Purchase(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  const marketAddress = props.marketAddress
  //var [user_id, setUserId] = useState(0)
  //var [user_name, setUserName] = useState<string>("")
  var [currency, setCurrency] = useState<string>("") // TODO: fix it to work as input option
  var [desCurrency, setDesCurrency] = useState<string>("")
  var [price,setPrice] = useState<string>("")
  var [human_number,setHuman_number] = useState<string>("")
  var [token_id,setTokenId] = useState(0)
  var [collection_address, setCollectionAddress] = useState<string>("")


  useEffect(() => {
    const queryParams = new URLSearchParams(location.search);


  
  var token_id_q = queryParams.get('token_id'); // amount MUST be in wei format
  let token_id_n : number = +token_id_q;
  setTokenId(token_id_n);

  
  var p = queryParams.get('price');
  console.log(p);
  var a_ether = ethers.utils.formatEther(p);
  console.log(a_ether);
  setHuman_number(a_ether);
  setPrice(p);

  /*
  var ac_q = queryParams.get('collection_contract_address'); 
  var ac_a = ethers.utils.getAddress(ac_q);
  setCollectionAddress(ac_a);
  */
  setCollectionAddress(addressContract)

  var c = queryParams.get('currency');
  setCurrency(c);
  
  const MetaMarketplaceRead:Contract = new ethers.Contract(marketAddress, abi)//todo
  var desiredPrice, desiredCurrency = MetaMarketplaceRead.getLastPrice(addressContract, token_id)
  console.log("Desired price:", desiredPrice)


  if (desiredPrice != undefined && desiredCurrency != undefined) {
    setDesCurrency(desiredCurrency)
  }
  
 

  
  }, []);
  

  async function purchase(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const user_address = signer._address
    const MetaMarketplace:Contract = new ethers.Contract(marketAddress, abi, signer)
    console.log("token id to interact raw: ", token_id)

    
    // var token_id_uint = ethers.utils.
    var currency_int = parseInt(currency)
    MetaMarketplace.purchase(collection_address,token_id,currency_int,price)
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("purchase receipt", receipt)})
        })
         .catch((e:Error) => console.log(e))
     }


  
  //const handleChange = (value:string) => setUserId(value)   
  
  return (
    <form onSubmit={purchase}>
    <FormControl>
      <FormLabel >Purchase an nft: </FormLabel>
      <Input id="token_id" type="number" required  onChange={(e) => setTokenId(parseInt(e.target.value))} value={token_id} my={3}/>
     
     <Input id="price" type="text" placeholder="bid price" required  onChange={(e) => setPrice(e.target.value)} value={price} my={3}/> 
     <Input id="collection_contract_address" type="text" required  onChange={(e) => setCollectionAddress(e.target.value)} value={collection_address} my={3}/>
      <div>
        <Text><b>Token id to buy</b>:{token_id}</Text>
        <Text><b>Currency</b>:{currency}</Text>
        <Text><b>Marketplace address</b>:{marketAddress}</Text>

    </div>
      <Button type="submit" isDisabled={!currentAccount || currency != desCurrency}>Buy</Button>
    </FormControl>
    </form>
  )
}