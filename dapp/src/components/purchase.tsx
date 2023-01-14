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
  var [desiredPrice, setDesiredPrice] = useState<string>("0")

  useEffect(() => {
    const queryParams = new URLSearchParams(location.search);


  
  var token_id_q = queryParams.get('token_id'); // amount MUST be in wei format
  let token_id_n : number = +token_id_q;
  setTokenId(token_id_n);

  
  var p = queryParams.get('price');
  setHuman_number(p);
  
  setPrice(p);
  //price = p;

  /*
  var ac_q = queryParams.get('collection_contract_address'); 
  var ac_a = ethers.utils.getAddress(ac_q);
  setCollectionAddress(ac_a);
  */
  if (addressContract != undefined || addressContract != "") {
    setCollectionAddress(addressContract);
  } else {
    var ac_q = queryParams.get('collection_contract_address'); 
    var ac_a = ethers.utils.getAddress(ac_q);
    setCollectionAddress(ac_a);
  }

  var c = queryParams.get('currency');
  
  setCurrency(c);
  //currency = c;
  const provider = new ethers.providers.JsonRpcProvider("https://polygon-mainnet-rpc.allthatnode.com:8545")
  
  //const signer = provider.getSigner()

  const MetaMarketplaceRead:Contract = new ethers.Contract(marketAddress, abi, provider)

  console.log("Market address:", marketAddress)

  

  
  
  const fetchData = async () => {

  for (let i = 0; desiredPrice == '0' && i < 6; i++ ) {
    var desiredPriceBN
    desiredPriceBN = await MetaMarketplaceRead.getFloorPriceByCurrency(addressContract, 4, i)
    
    setDesiredPrice(desiredPriceBN.toString())

    setDesCurrency(i.toString())
    
    console.log("Desired price:", desiredPrice)
    console.log("Desired currency: ", desCurrency);

    console.log("Actual price:", price)
    console.log("Actual currency:", currency)
  
    

  }

}

fetchData()


  

  // TODO: add getter for active sell offer in contract?
  /*
  var MarketplaceStruct = MetaMarketplaceRead.Marketplaces(collection_address);
  console.log("Try to get Marketplace struct: ", MarketplaceStruct);
  */



 
  


  
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

function countDecimals(pri, cur) {
  var mindiv
  console.log("unedited price is:",pri)
  console.log("unedited amount to pay in wei: ", price)
    if (pri == "") {
      pri = 0
    }
    if (cur == '0' || cur == '1') { //usdt
      mindiv = 6
    } else  if (addressContract == '4') { //wbtc
      mindiv = 8
    } else  {
      mindiv = 18
    }
    
    
    let amount_wei = ethers.utils.parseUnits(pri, mindiv)  // let's suppose we got wei number in query and we just parse it
    console.log("amount to pay in wei: ", price)
    console.log("currency", currency)
    return amount_wei.toString()
}
  
  //const handleChange = (value:string) => setUserId(value)   
  
  return (
    <form onSubmit={purchase}>
    <FormControl>
      <FormLabel >Token id & token contract address: </FormLabel>
      <Input id="token_id" type="number" placeholder="Token ID"  required  onChange={(e) => setTokenId(parseInt(e.target.value))} value={token_id} my={3}/>
      <Input id="collection_contract_address" type="text" required  onChange={(e) => setCollectionAddress(e.target.value)} value={collection_address} my={3}/>
    
     
     
      <div>
      <Text>Select amount & currency in which you wish to pay:</Text>
      <Input id="price" type="text" placeholder="Bid amount" required  onChange={(e) => setPrice(countDecimals(e.target.value, currency))} onInput={(e) => setHuman_number(e.currentTarget.value)} value={human_number} my={3}/> 
      <Select id="currency" placeholder="Select currency in which you want to pay:" onChange={(e) => setCurrency(e.target.value)} value= {currency} onInput={(e) => setPrice(countDecimals(human_number, e.currentTarget.value))} my={3}>
      <option value='0'>USDT</option>
      <option value='1'>USDC</option>
      <option value='2'>DAI</option>
      <option value='3'>WETH</option>
      <option value='4'>WBTC</option>
      <option value='5'>VXPPL</option>
      </Select>
        <Text><b>Marketplace address</b>:{marketAddress}</Text>

    </div>
      <Button type="submit" isDisabled={price != desiredPrice || currency != desCurrency}>Buy</Button>
    </FormControl>
    </form>
  )
}