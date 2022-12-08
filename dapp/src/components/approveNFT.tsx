import React, { useState, useEffect } from 'react'
import {Button, Input , NumberInput,  NumberInputField,  FormControl,  FormLabel, Text } from '@chakra-ui/react'
import {ethers} from 'ethers'
import {parseEther } from 'ethers/lib/utils'
import {abi} from '../../../artifacts/@openzeppelin/contracts/token/ERC721/ERC721.sol/ERC721.json'
import { Contract } from "ethers"
import { TransactionResponse,TransactionReceipt } from "@ethersproject/abstract-provider"


interface Props {
   // addressContract: string,
    currentAccount: string | undefined,
    marketAddress: string
}

declare let window: any;


export default function ApproveSingleton(props:Props){
  //const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  const marketAddress = props.marketAddress
  //var [user_id, setUserId] = useState(0)
  //var [user_name, setUserName] = useState<string>("")
  //var [amount,setAmount] = useState<string>("")
  //var [human_number,setHuman_number] = useState<string>("")
  var [token_id,setTokenId] = useState(0)
  var [addressContract, setAddressContract] = useState<string>("")

  useEffect(() => {
    const queryParams = new URLSearchParams(location.search);


  
  var token_id_q = queryParams.get('token_id'); // amount MUST be in wei format
  let token_id_n : number = +token_id_q;
  setTokenId(token_id_n);
  
  var ac_q = queryParams.get('token_contract_address'); // erc20 to approve
  var ac_a = ethers.utils.getAddress(ac_q);
  setAddressContract(ac_a);
  
  }, []);
  

  async function approveNft(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const user_address = signer._address
    const NFT_Contract:Contract = new ethers.Contract(addressContract, abi, signer)
    console.log("token id to approve raw: ", token_id)
    // var token_id_uint = ethers.utils.
 
   // let passport_fee_wei = ethers.utils.formatUnits(1000,"wei");
    //let passport_fee_custom_gwei = ethers.utils.formatUnits(2000000,"gwei"); // 1 gwei = 1'000'000'000 wei, 2m gwei = 0,002 (estimateGas on approval = 0.02, so we need to take that fee for gas)
    //let passport_fee_wei = ethers.utils.formatUnits(passport_fee_custom_gwei,"wei");
    //let passport_fee_wei_hardcode = ethers.utils.formatUnits(2000000000000000,"wei");
    NFT_Contract.approve(marketAddress,token_id)
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("approve receipt", receipt)})
        })
         .catch((e:Error) => console.log(e))
     }


  
  //const handleChange = (value:string) => setUserId(value)
  
  return (
    <form onSubmit={approveNft}>
    <FormControl>
      <FormLabel >Approve NFT to marketplace: </FormLabel>
      <div>
        <Text><b>Token Contract</b>: {addressContract}</Text>
        <Text><b>Token id to approve</b>:{token_id}</Text>
        <Text><b>Marketplace address</b>:{marketAddress}</Text>
    </div>
      <Button type="submit" isDisabled={!currentAccount}>Approve</Button>
    </FormControl>
    </form>
  )
}