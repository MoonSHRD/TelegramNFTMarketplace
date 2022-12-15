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


export default function SetupMarketplace(props:Props){
  const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  const marketAddress = props.marketAddress
  //var [user_id, setUserId] = useState(0)
  //var [user_name, setUserName] = useState<string>("")
 // var [currency, setCurrency] = useState(0) // TODO: fix it to work as input option
 // var [price,setPrice] = useState<string>("")
  var [human_number,setHuman_number] = useState<string>("")
 // var [token_id,setTokenId] = useState(0)
  var [collection_address, setCollectionAddress] = useState<string>("")
  var [owner_address, setOwnerAddress] = useState<string>("")
  var [fee, setFee] = useState<string>("")
  var [n_type, setNType] = useState<string>("")


  useEffect(() => {
  const queryParams = new URLSearchParams(location.search);

    if(!window.ethereum) return

  const provider = new ethers.providers.Web3Provider(window.ethereum)
  
 // var owner_q = queryParams.get('owner');
  setOwnerAddress(currentAccount);

  var p = queryParams.get('fee');
  var a_ether = ethers.utils.formatEther(p);
  setHuman_number(a_ether);
  setFee(p);

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
  

  async function setupMarketplace(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const user_address = signer._address
    const MetaMarketplace:Contract = new ethers.Contract(addressContract, abi, signer)
    //console.log("token id to interact raw: ", token_id)
    // var token_id_uint = ethers.utils.
 
    MetaMarketplace.SetUpMarketplace(collection_address,n_type,owner_address,fee)
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("setup marketplace receipt", receipt)})
        })
         .catch((e:Error) => console.log(e))
     }


  
  //const handleChange = (value:string) => setUserId(value)   
  
  return (
    <form onSubmit={setupMarketplace}>
    <FormControl>
      <FormLabel >Setup your collection: </FormLabel> 
     <Input id="collection_contract_address" type="text" required  onChange={(e) => setCollectionAddress(e.target.value)} value={collection_address} my={3}/>
     <Input id="fee" type="number" required  onChange={(e) => setFee(e.target.value)} value={fee} my={3}/>
     <Select id="n_type" placeholder="Select NFT type:" onChange={(e) => setNType(e.target.value)} value= {n_type}  my={3}>
      <option value='1'>Enum</option>
      <option value='2'>Meta</option>
      <option value='3'>Common</option>
      <option value='4'>URIStorage</option>
      </Select>
      <div>
        <Text><b>Owner</b>:{currentAccount}</Text>

    </div>
      <Button type="submit" isDisabled={!currentAccount}>Setup</Button>
    </FormControl>
    </form>
  )
}