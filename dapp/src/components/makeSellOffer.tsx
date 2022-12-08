import React, { useState, useEffect } from 'react'
import {Button, Input , NumberInput,  NumberInputField,  FormControl,  FormLabel, Text } from '@chakra-ui/react'
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


export default function MakeSellOffer(props:Props){
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


  
  var token_id_q = queryParams.get('token_id'); // amount MUST be in wei format
  let token_id_n : number = +token_id_q;
  setTokenId(token_id_n);

  
  var p = queryParams.get('min_price');
  var a_ether = ethers.utils.formatEther(p);
  setHuman_number(a_ether);
  

  var ac_q = queryParams.get('collection_contract_address'); // erc20 to approve
  var ac_a = ethers.utils.getAddress(ac_q);
  setCollectionAddress(ac_a);

  
  
  
  
  }, []);
  

  async function approveSingleton(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const user_address = signer._address
    const MetaMarketplace:Contract = new ethers.Contract(addressContract, abi, signer)
    console.log("token id to interact raw: ", token_id)
    // var token_id_uint = ethers.utils.
 
   // let passport_fee_wei = ethers.utils.formatUnits(1000,"wei");
    //let passport_fee_custom_gwei = ethers.utils.formatUnits(2000000,"gwei"); // 1 gwei = 1'000'000'000 wei, 2m gwei = 0,002 (estimateGas on approval = 0.02, so we need to take that fee for gas)
    //let passport_fee_wei = ethers.utils.formatUnits(passport_fee_custom_gwei,"wei");
    //let passport_fee_wei_hardcode = ethers.utils.formatUnits(2000000000000000,"wei");
    MetaMarketplace.makeSellOffer(token_id,price,collection_address,currency)
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("make sell offer receipt", receipt)})
        })
         .catch((e:Error) => console.log(e))
     }


  
  //const handleChange = (value:string) => setUserId(value)   setCurrency
  
  return (
    <form onSubmit={approveSingleton}>
    <FormControl>
      <FormLabel >Sell your nft: </FormLabel>
      <Input id="token_id" type="number" required  onChange={(e) => setTokenId(parseInt(e.target.value))} value={token_id} my={3}/>
     
     <Input id="price" type="text" required  onChange={(e) => setPrice(e.target.value)} value={price} my={3}/> 
     <Input id="collection_contract_address" type="text" required  onChange={(e) => setCollectionAddress(e.target.value)} value={collection_address} my={3}/>
     <div>
     <Input id="currency_USDT" type="radio" required  name="currency_radio" onChange={(e) => setCurrency(e.target.value)} value={0}   />
     <label for="currency_USDT">USDT</label>
     <Input id="currency_USDC" type="radio"  name="currency_radio" onChange={(e) => setCurrency(e.target.value)} value={1}   />
     <label for="currency_USDC">USDC</label>
     <Input id="currency_DAI" type="radio"  name="currency_radio" onChange={(e) => setCurrency(e.target.value)} value={2}   />
     <label for="currency_DAI">DAI</label>
     <Input id="currency_WETH" type="radio"  name="currency_radio" onChange={(e) => setCurrency(e.target.value)} value={3}   />
     <label for="currency_WETH">WETH</label>
     <Input id="currency_WBTC" type="radio"  name="currency_radio" onChange={(e) => setCurrency(e.target.value)} value={4}   />
     <label for="currency_WBTC">WBTC</label>
     <Input id="currency_VXPPL" type="radio"  name="currency_radio" onChange={(e) => setCurrency(e.target.value)} value={5}   />
     <label for="currency_VXPPL">VXPPL</label>
     </div>
      <div>
        <Text><b>Token id to approve</b>:{token_id}</Text>
        <Text><b>Marketplace address</b>:{addressContract}</Text>
    </div>
      <Button type="submit" isDisabled={!currentAccount}>Approve</Button>
    </FormControl>
    </form>
  )
}