import React, { useState, useEffect } from 'react'
import {Button, Input , NumberInput,  NumberInputField, Select, FormControl,  FormLabel, Text } from '@chakra-ui/react'
import {ethers} from 'ethers'
import {parseEther } from 'ethers/lib/utils'
import {abi} from '../../../artifacts/@openzeppelin/contracts/token/ERC20/IERC20.sol/IERC20.json'
import { Contract } from "ethers"
import { TransactionResponse,TransactionReceipt } from "@ethersproject/abstract-provider"
import {toWei} from '../../../scripts/lib.js'

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

  var [human_number,setHuman_number] = useState<string>("")
  var [addressContract, setAddressContract] = useState<string>("")


  useEffect(() => {
    

  }, []);
  

  async function approveERC20(event:React.FormEvent) {
    console.log(addressContract)
    console.log(marketAddress)
    event.preventDefault()
    if(!window.ethereum) return    
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    const user_address = currentAccount
    console.log(user_address)
    const ERC20_contract:Contract = new ethers.Contract(addressContract, abi, signer)
    
    var mindiv

    if (addressContract == '0xc2132D05D31c914a87C6611C10748AEb04B58e8F' || addressContract == '0x2791bca1f2de4661ed88a30c99a7a9449aa84174') {
      mindiv = 6
    } else  if (addressContract == '0x1BFD67037B42Cf73acF2047067bd4F2C47D9BfD6') {
      mindiv = 8
    } else  {
      mindiv = 18
    }




    let amount_wei = ethers.utils.parseUnits(human_number.toString(), mindiv)  // let's suppose we got wei number in query and we just parse it
    console.log("amount to approve in wei: ", amount_wei)
   // let passport_fee_wei = ethers.utils.formatUnits(1000,"wei");
    //let passport_fee_custom_gwei = ethers.utils.formatUnits(2000000,"gwei"); // 1 gwei = 1'000'000'000 wei, 2m gwei = 0,002 (estimateGas on approval = 0.02, so we need to take that fee for gas)
    //let passport_fee_wei = ethers.utils.formatUnits(passport_fee_custom_gwei,"wei");
    //let passport_fee_wei_hardcode = ethers.utils.formatUnits(2000000000000000,"wei");
    ERC20_contract.approve(marketAddress, amount_wei, {from: user_address})
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
      <div>
      <Input id="amount" type="text" placeholder="Enter amount of tokens to approve" required  onChange={(e) => setHuman_number(e.target.value)} value={human_number} my={3}/>
     <Select id="currency" placeholder="Select currency you want to approve:" onChange={(e) => setAddressContract(e.target.value)} value= {addressContract}  my={3}>
      <option value='0xc2132D05D31c914a87C6611C10748AEb04B58e8F'>USDT</option>
      <option value='0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174'>USDC</option>
      <option value='0x8f3Cf7ad23Cd3CaDbD9735AFf958023239c6A063'>DAI</option>
      <option value='0x7ceb23fd6bc0add59e62ac25578270cff1b9f619'>WETH</option>
      <option value='0x1BFD67037B42Cf73acF2047067bd4F2C47D9BfD6'>WBTC</option>
      <option value='0xA25d4cb14742a408B7421A259c96cE29F6113B3d'>VXPPL</option>
      </Select>
     </div>
    </div>
      <Button type="submit" isDisabled={!currentAccount}>Approve</Button>
    </FormControl>
    </form>
  )
}