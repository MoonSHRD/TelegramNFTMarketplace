import React, { useState, useEffect } from 'react'
import {Button, Input , NumberInput,  NumberInputField,  FormControl,  FormLabel, Text } from '@chakra-ui/react'
import {ethers} from 'ethers'
import {parseEther } from 'ethers/lib/utils'
import {abi} from '../../../artifacts/@openzeppelin/contracts/token/ERC20/IERC20.sol/IERC20.json'
import { Contract } from "ethers"
import { TransactionResponse,TransactionReceipt } from "@ethersproject/abstract-provider"


interface Props {
   // addressContract: string,
    currentAccount: string | undefined ,
    marketAddress: string
}

declare let window: any;


export default function ApproveERC20(props:Props){
 // const addressContract = props.addressContract
  const currentAccount = props.currentAccount
  const marketAddress = props.marketAddress     // market address
  //var [user_id, setUserId] = useState(0)
  //var [user_name, setUserName] = useState<string>("")
  var [amount,setAmount] = useState<string>("")
  var [human_number,setHuman_number] = useState<string>("")
  var [addressContract, setAddressContract] = useState<string>("")


  useEffect(() => {
    const queryParams = new URLSearchParams(location.search);


  
  var amount_q = queryParams.get('amount'); // amount MUST be in wei format
  //let amount_n : number = +amount_q;
  setAmount(amount_q);
  var a_wei = ethers.utils.parseUnits(amount_q,"wei");
  var a_ether = ethers.utils.formatEther(a_wei);
  setHuman_number(a_ether);

  var ac_q = queryParams.get('token_contract_address'); // erc20 to approve
  var ac_a = ethers.utils.getAddress(ac_q);
  setAddressContract(ac_a);
  
  }, []);
  

  async function approveERC20(event:React.FormEvent) {
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const user_address = signer._address
    const ERC20_contract:Contract = new ethers.Contract(addressContract, abi, signer)
    console.log("amount to approve raw: ", amount)
    var amount_wei = ethers.utils.parseUnits(amount,"wei")  // let's suppose we got wei number in query and we just parse it
    console.log("amount to approve in wei: ", amount_wei)
    var amount_ether = ethers.utils.formatEther(amount_wei)
    console.log ("amount in Coin format: ", amount_ether)
   // let passport_fee_wei = ethers.utils.formatUnits(1000,"wei");
    //let passport_fee_custom_gwei = ethers.utils.formatUnits(2000000,"gwei"); // 1 gwei = 1'000'000'000 wei, 2m gwei = 0,002 (estimateGas on approval = 0.02, so we need to take that fee for gas)
    //let passport_fee_wei = ethers.utils.formatUnits(passport_fee_custom_gwei,"wei");
    //let passport_fee_wei_hardcode = ethers.utils.formatUnits(2000000000000000,"wei");
    ERC20_contract.approve(user_address,marketAddress,{value:amount_wei})
     .then((tr: TransactionResponse) => {
        console.log(`TransactionResponse TX hash: ${tr.hash}`)
        tr.wait().then((receipt:TransactionReceipt) => {console.log("approve receipt", receipt)})
        })
         .catch((e:Error) => console.log(e))
     }


  
  //const handleChange = (value:string) => setUserId(value)
  
  return (
    <form onSubmit={approveERC20}>
    <FormControl>
      <FormLabel >Approve ERC20 to marketplace: </FormLabel>
      <div>
        <Text><b>Token Contract</b>: {addressContract}</Text>
        <Text><b>Amount of token to approve</b>:{human_number}</Text>
        <Text><b>Marketplace address</b>:{marketAddress}</Text>
    </div>
      <Button type="submit" isDisabled={!currentAccount}>Approve</Button>
    </FormControl>
    </form>
  )
}