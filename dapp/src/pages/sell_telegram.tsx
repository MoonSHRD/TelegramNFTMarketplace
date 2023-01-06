import type { NextPage } from 'next'
import Head from 'next/head'
import NextLink from "next/link"
import { VStack, Heading, Box, LinkOverlay, LinkBox} from "@chakra-ui/layout"
import { Text, Button } from '@chakra-ui/react'
import { useState, useEffect} from 'react'
import {ethers} from "ethers"
import ReadNFT_Telegram from "../components/readNFT_Telegram"
import ApproveSingleton from "../components/approveSingletonNft"
import MakeSellOffer from "../components/makeSellOffer"


//buy_telegram?token_id=1&currency=0&price=1&collection_contract_address=0xEbE648689E98abA446e38621E5a3491db03a7621

declare let window:any

const Home: NextPage = () => {
  const [balance, setBalance] = useState<string | undefined>()
  const [currentAccount, setCurrentAccount] = useState<string | undefined>()
  const [chainId, setChainId] = useState<number | undefined>()
  const [chainname, setChainName] = useState<string | undefined>()

  useEffect(() => {
    if(!currentAccount || !ethers.utils.isAddress(currentAccount)) return
    //client side code
    if(!window.ethereum) return
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    provider.getBalance(currentAccount).then((result)=>{
      setBalance(ethers.utils.formatEther(result))
    })
    provider.getNetwork().then((result)=>{
      setChainId(result.chainId)
      setChainName(result.name)
    })

  },[currentAccount])

  const onClickConnect = () => {
    //client side code
    if(!window.ethereum) {
      console.log("please install MetaMask")
      return
    }
    /*
    //change from window.ethereum.enable() which is deprecated
    //see docs: https://docs.metamask.io/guide/ethereum-provider.html#legacy-methods
    window.ethereum.request({ method: 'eth_requestAccounts' })
    .then((accounts:any)=>{
      if(accounts.length>0) setCurrentAccount(accounts[0])
    })
    .catch('error',console.error)
    */

    //we can do it using ethers.js
    const provider = new ethers.providers.Web3Provider(window.ethereum)

    // MetaMask requires requesting permission to connect users accounts
    provider.send("eth_requestAccounts", [])
    .then((accounts)=>{
      if(accounts.length>0) setCurrentAccount(accounts[0])
    })
    .catch((e)=>console.log(e))
  }

  const onClickDisconnect = () => {
    console.log("onClickDisConnect")
    setBalance(undefined)
    setCurrentAccount(undefined)
  }

  return (
    <>
      <Head>
        <title>NFT Marketplace</title>
      </Head>

      <Heading as="h3"  my={4}>Sell your NFT</Heading>          
      <VStack>
        <Box w='100%' my={4}>
        {currentAccount  
          ? <Button type="button" w='100%' onClick={onClickDisconnect}>
                Account:{currentAccount}
            </Button>
          : <Button type="button" w='100%' onClick={onClickConnect}>
                  Connect MetaMask
              </Button>
        }
        </Box>
        {currentAccount  
          ?<Box  mb={0} p={4} w='100%' borderWidth="1px" borderRadius="lg">
          <Heading my={4}  fontSize='xl'>Account info</Heading>
          <Text>ETH Balance of current account: {balance}</Text>
          <Text>Chain Info: ChainId {chainId} name {chainname}</Text>
          <Text><b>NETWORK SHOULD BE POLYGON!</b></Text>
        </Box>
        :<></>
        }

        <Box  mb={0} p={4} w='100%' borderWidth="1px" borderRadius="lg">
          <Heading my={4}  fontSize='xl'>Your NFT:</Heading>
          <ReadNFT_Telegram 
            addressContract='0xEbE648689E98abA446e38621E5a3491db03a7621'
            currentAccount={currentAccount}
          />
        </Box> 
        <Box  mb={0} p={4} w='100%' borderWidth="1px" borderRadius="lg">
          <Heading my={4}  fontSize='xl'>Approve your NFT to be interactable at marketplace:</Heading>
          <ApproveSingleton 
            addressContract='0xEbE648689E98abA446e38621E5a3491db03a7621'
            marketAddress='0x8CeC1dD2802C820574b34357b4AdD49aaF100fAD'
            currentAccount={currentAccount}
          />
        </Box> 
        <Box  mb={0} p={4} w='100%' borderWidth="1px" borderRadius="lg">
          <Heading my={4}  fontSize='xl'>Put your NFT on sale</Heading>
          <MakeSellOffer 
            addressContract='0x8CeC1dD2802C820574b34357b4AdD49aaF100fAD'
            marketAddress='0x8CeC1dD2802C820574b34357b4AdD49aaF100fAD'
            collectionContract='0xEbE648689E98abA446e38621E5a3491db03a7621'
            currentAccount={currentAccount}
          />
        </Box> 
  

...
      </VStack>
    </>
  )
}

export default Home