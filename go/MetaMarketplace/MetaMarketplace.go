// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MetaMarketplace

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MetaMarketplaceMetaData contains all meta data concerning the MetaMarketplace contract.
var MetaMarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currency_contract_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"telegram_collection_\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"treasure_fund_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"BuyOfferWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initial_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transfered_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeAddress\",\"type\":\"address\"}],\"name\":\"CalculatedFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"NewBuyOffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft_contract_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"NewSellOffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft_contract_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recepient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RoyaltiesPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft_contract_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Sale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft_contract_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"SellOfferWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Marketplaces\",\"outputs\":[{\"internalType\":\"enumMetaMarketplace.NftType\",\"name\":\"nft_standard\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"collectionOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ownerFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"promille_fee_\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"treasure_fund_\",\"type\":\"address\"}],\"name\":\"SetServiceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft_contract_\",\"type\":\"address\"},{\"internalType\":\"enumMetaMarketplace.NftType\",\"name\":\"standard_\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"collection_owner_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collection_fee_\",\"type\":\"uint256\"}],\"name\":\"SetUpMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_treasure_fund\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_contract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumCurrenciesERC20.CurrencyERC20\",\"name\":\"currency_\",\"type\":\"uint8\"}],\"name\":\"acceptBuyOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"promille_fee_\",\"type\":\"uint256\"}],\"name\":\"calculateAbstractFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft_contract_\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"collection_owner_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collection_fee_\",\"type\":\"uint256\"}],\"name\":\"editMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_contract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getLastPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_lastPrice\",\"type\":\"uint256\"},{\"internalType\":\"enumCurrenciesERC20.CurrencyERC20\",\"name\":\"currency_\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_contract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumCurrenciesERC20.CurrencyERC20\",\"name\":\"currency_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"bid_price_\",\"type\":\"uint256\"}],\"name\":\"makeBuyOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nft_contract_\",\"type\":\"address\"},{\"internalType\":\"enumCurrenciesERC20.CurrencyERC20\",\"name\":\"currency_\",\"type\":\"uint8\"}],\"name\":\"makeSellOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"promille_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_contract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumCurrenciesERC20.CurrencyERC20\",\"name\":\"currency_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"bid_price_\",\"type\":\"uint256\"}],\"name\":\"purchase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_contract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumCurrenciesERC20.CurrencyERC20\",\"name\":\"currency_\",\"type\":\"uint8\"}],\"name\":\"withdrawBuyOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft_contract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawSellOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MetaMarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MetaMarketplaceMetaData.ABI instead.
var MetaMarketplaceABI = MetaMarketplaceMetaData.ABI

// MetaMarketplace is an auto generated Go binding around an Ethereum contract.
type MetaMarketplace struct {
	MetaMarketplaceCaller     // Read-only binding to the contract
	MetaMarketplaceTransactor // Write-only binding to the contract
	MetaMarketplaceFilterer   // Log filterer for contract events
}

// MetaMarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetaMarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaMarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetaMarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaMarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetaMarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaMarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetaMarketplaceSession struct {
	Contract     *MetaMarketplace  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetaMarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetaMarketplaceCallerSession struct {
	Contract *MetaMarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MetaMarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetaMarketplaceTransactorSession struct {
	Contract     *MetaMarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MetaMarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetaMarketplaceRaw struct {
	Contract *MetaMarketplace // Generic contract binding to access the raw methods on
}

// MetaMarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetaMarketplaceCallerRaw struct {
	Contract *MetaMarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MetaMarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetaMarketplaceTransactorRaw struct {
	Contract *MetaMarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetaMarketplace creates a new instance of MetaMarketplace, bound to a specific deployed contract.
func NewMetaMarketplace(address common.Address, backend bind.ContractBackend) (*MetaMarketplace, error) {
	contract, err := bindMetaMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetaMarketplace{MetaMarketplaceCaller: MetaMarketplaceCaller{contract: contract}, MetaMarketplaceTransactor: MetaMarketplaceTransactor{contract: contract}, MetaMarketplaceFilterer: MetaMarketplaceFilterer{contract: contract}}, nil
}

// NewMetaMarketplaceCaller creates a new read-only instance of MetaMarketplace, bound to a specific deployed contract.
func NewMetaMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MetaMarketplaceCaller, error) {
	contract, err := bindMetaMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaMarketplaceCaller{contract: contract}, nil
}

// NewMetaMarketplaceTransactor creates a new write-only instance of MetaMarketplace, bound to a specific deployed contract.
func NewMetaMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MetaMarketplaceTransactor, error) {
	contract, err := bindMetaMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaMarketplaceTransactor{contract: contract}, nil
}

// NewMetaMarketplaceFilterer creates a new log filterer instance of MetaMarketplace, bound to a specific deployed contract.
func NewMetaMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MetaMarketplaceFilterer, error) {
	contract, err := bindMetaMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetaMarketplaceFilterer{contract: contract}, nil
}

// bindMetaMarketplace binds a generic wrapper to an already deployed contract.
func bindMetaMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaMarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaMarketplace *MetaMarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaMarketplace.Contract.MetaMarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaMarketplace *MetaMarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.MetaMarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaMarketplace *MetaMarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.MetaMarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaMarketplace *MetaMarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaMarketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaMarketplace *MetaMarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaMarketplace *MetaMarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.contract.Transact(opts, method, params...)
}

// Marketplaces is a free data retrieval call binding the contract method 0xfa75c8c5.
//
// Solidity: function Marketplaces(address ) view returns(uint8 nft_standard, bool initialized, address collectionOwner, uint256 ownerFee)
func (_MetaMarketplace *MetaMarketplaceCaller) Marketplaces(opts *bind.CallOpts, arg0 common.Address) (struct {
	NftStandard     uint8
	Initialized     bool
	CollectionOwner common.Address
	OwnerFee        *big.Int
}, error) {
	var out []interface{}
	err := _MetaMarketplace.contract.Call(opts, &out, "Marketplaces", arg0)

	outstruct := new(struct {
		NftStandard     uint8
		Initialized     bool
		CollectionOwner common.Address
		OwnerFee        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NftStandard = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Initialized = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.CollectionOwner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.OwnerFee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Marketplaces is a free data retrieval call binding the contract method 0xfa75c8c5.
//
// Solidity: function Marketplaces(address ) view returns(uint8 nft_standard, bool initialized, address collectionOwner, uint256 ownerFee)
func (_MetaMarketplace *MetaMarketplaceSession) Marketplaces(arg0 common.Address) (struct {
	NftStandard     uint8
	Initialized     bool
	CollectionOwner common.Address
	OwnerFee        *big.Int
}, error) {
	return _MetaMarketplace.Contract.Marketplaces(&_MetaMarketplace.CallOpts, arg0)
}

// Marketplaces is a free data retrieval call binding the contract method 0xfa75c8c5.
//
// Solidity: function Marketplaces(address ) view returns(uint8 nft_standard, bool initialized, address collectionOwner, uint256 ownerFee)
func (_MetaMarketplace *MetaMarketplaceCallerSession) Marketplaces(arg0 common.Address) (struct {
	NftStandard     uint8
	Initialized     bool
	CollectionOwner common.Address
	OwnerFee        *big.Int
}, error) {
	return _MetaMarketplace.Contract.Marketplaces(&_MetaMarketplace.CallOpts, arg0)
}

// TreasureFund is a free data retrieval call binding the contract method 0x586e5576.
//
// Solidity: function _treasure_fund() view returns(address)
func (_MetaMarketplace *MetaMarketplaceCaller) TreasureFund(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaMarketplace.contract.Call(opts, &out, "_treasure_fund")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasureFund is a free data retrieval call binding the contract method 0x586e5576.
//
// Solidity: function _treasure_fund() view returns(address)
func (_MetaMarketplace *MetaMarketplaceSession) TreasureFund() (common.Address, error) {
	return _MetaMarketplace.Contract.TreasureFund(&_MetaMarketplace.CallOpts)
}

// TreasureFund is a free data retrieval call binding the contract method 0x586e5576.
//
// Solidity: function _treasure_fund() view returns(address)
func (_MetaMarketplace *MetaMarketplaceCallerSession) TreasureFund() (common.Address, error) {
	return _MetaMarketplace.Contract.TreasureFund(&_MetaMarketplace.CallOpts)
}

// CalculateAbstractFee is a free data retrieval call binding the contract method 0x36ea6178.
//
// Solidity: function calculateAbstractFee(uint256 amount, uint256 scale, uint256 promille_fee_) pure returns(uint256)
func (_MetaMarketplace *MetaMarketplaceCaller) CalculateAbstractFee(opts *bind.CallOpts, amount *big.Int, scale *big.Int, promille_fee_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetaMarketplace.contract.Call(opts, &out, "calculateAbstractFee", amount, scale, promille_fee_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateAbstractFee is a free data retrieval call binding the contract method 0x36ea6178.
//
// Solidity: function calculateAbstractFee(uint256 amount, uint256 scale, uint256 promille_fee_) pure returns(uint256)
func (_MetaMarketplace *MetaMarketplaceSession) CalculateAbstractFee(amount *big.Int, scale *big.Int, promille_fee_ *big.Int) (*big.Int, error) {
	return _MetaMarketplace.Contract.CalculateAbstractFee(&_MetaMarketplace.CallOpts, amount, scale, promille_fee_)
}

// CalculateAbstractFee is a free data retrieval call binding the contract method 0x36ea6178.
//
// Solidity: function calculateAbstractFee(uint256 amount, uint256 scale, uint256 promille_fee_) pure returns(uint256)
func (_MetaMarketplace *MetaMarketplaceCallerSession) CalculateAbstractFee(amount *big.Int, scale *big.Int, promille_fee_ *big.Int) (*big.Int, error) {
	return _MetaMarketplace.Contract.CalculateAbstractFee(&_MetaMarketplace.CallOpts, amount, scale, promille_fee_)
}

// GetLastPrice is a free data retrieval call binding the contract method 0x1922cb8d.
//
// Solidity: function getLastPrice(address token_contract_, uint256 _tokenId) view returns(uint256 _lastPrice, uint8 currency_)
func (_MetaMarketplace *MetaMarketplaceCaller) GetLastPrice(opts *bind.CallOpts, token_contract_ common.Address, _tokenId *big.Int) (struct {
	LastPrice *big.Int
	Currency  uint8
}, error) {
	var out []interface{}
	err := _MetaMarketplace.contract.Call(opts, &out, "getLastPrice", token_contract_, _tokenId)

	outstruct := new(struct {
		LastPrice *big.Int
		Currency  uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastPrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Currency = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetLastPrice is a free data retrieval call binding the contract method 0x1922cb8d.
//
// Solidity: function getLastPrice(address token_contract_, uint256 _tokenId) view returns(uint256 _lastPrice, uint8 currency_)
func (_MetaMarketplace *MetaMarketplaceSession) GetLastPrice(token_contract_ common.Address, _tokenId *big.Int) (struct {
	LastPrice *big.Int
	Currency  uint8
}, error) {
	return _MetaMarketplace.Contract.GetLastPrice(&_MetaMarketplace.CallOpts, token_contract_, _tokenId)
}

// GetLastPrice is a free data retrieval call binding the contract method 0x1922cb8d.
//
// Solidity: function getLastPrice(address token_contract_, uint256 _tokenId) view returns(uint256 _lastPrice, uint8 currency_)
func (_MetaMarketplace *MetaMarketplaceCallerSession) GetLastPrice(token_contract_ common.Address, _tokenId *big.Int) (struct {
	LastPrice *big.Int
	Currency  uint8
}, error) {
	return _MetaMarketplace.Contract.GetLastPrice(&_MetaMarketplace.CallOpts, token_contract_, _tokenId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetaMarketplace *MetaMarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaMarketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetaMarketplace *MetaMarketplaceSession) Owner() (common.Address, error) {
	return _MetaMarketplace.Contract.Owner(&_MetaMarketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetaMarketplace *MetaMarketplaceCallerSession) Owner() (common.Address, error) {
	return _MetaMarketplace.Contract.Owner(&_MetaMarketplace.CallOpts)
}

// PromilleFee is a free data retrieval call binding the contract method 0x3d56bb9f.
//
// Solidity: function promille_fee() view returns(uint256)
func (_MetaMarketplace *MetaMarketplaceCaller) PromilleFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaMarketplace.contract.Call(opts, &out, "promille_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PromilleFee is a free data retrieval call binding the contract method 0x3d56bb9f.
//
// Solidity: function promille_fee() view returns(uint256)
func (_MetaMarketplace *MetaMarketplaceSession) PromilleFee() (*big.Int, error) {
	return _MetaMarketplace.Contract.PromilleFee(&_MetaMarketplace.CallOpts)
}

// PromilleFee is a free data retrieval call binding the contract method 0x3d56bb9f.
//
// Solidity: function promille_fee() view returns(uint256)
func (_MetaMarketplace *MetaMarketplaceCallerSession) PromilleFee() (*big.Int, error) {
	return _MetaMarketplace.Contract.PromilleFee(&_MetaMarketplace.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetaMarketplace *MetaMarketplaceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MetaMarketplace.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetaMarketplace *MetaMarketplaceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MetaMarketplace.Contract.SupportsInterface(&_MetaMarketplace.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetaMarketplace *MetaMarketplaceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MetaMarketplace.Contract.SupportsInterface(&_MetaMarketplace.CallOpts, interfaceId)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x505adff4.
//
// Solidity: function SetServiceFee(uint256 promille_fee_, address treasure_fund_) returns()
func (_MetaMarketplace *MetaMarketplaceTransactor) SetServiceFee(opts *bind.TransactOpts, promille_fee_ *big.Int, treasure_fund_ common.Address) (*types.Transaction, error) {
	return _MetaMarketplace.contract.Transact(opts, "SetServiceFee", promille_fee_, treasure_fund_)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x505adff4.
//
// Solidity: function SetServiceFee(uint256 promille_fee_, address treasure_fund_) returns()
func (_MetaMarketplace *MetaMarketplaceSession) SetServiceFee(promille_fee_ *big.Int, treasure_fund_ common.Address) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.SetServiceFee(&_MetaMarketplace.TransactOpts, promille_fee_, treasure_fund_)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x505adff4.
//
// Solidity: function SetServiceFee(uint256 promille_fee_, address treasure_fund_) returns()
func (_MetaMarketplace *MetaMarketplaceTransactorSession) SetServiceFee(promille_fee_ *big.Int, treasure_fund_ common.Address) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.SetServiceFee(&_MetaMarketplace.TransactOpts, promille_fee_, treasure_fund_)
}

// SetUpMarketplace is a paid mutator transaction binding the contract method 0xf82ffff2.
//
// Solidity: function SetUpMarketplace(address nft_contract_, uint8 standard_, address collection_owner_, uint256 collection_fee_) returns()
func (_MetaMarketplace *MetaMarketplaceTransactor) SetUpMarketplace(opts *bind.TransactOpts, nft_contract_ common.Address, standard_ uint8, collection_owner_ common.Address, collection_fee_ *big.Int) (*types.Transaction, error) {
	return _MetaMarketplace.contract.Transact(opts, "SetUpMarketplace", nft_contract_, standard_, collection_owner_, collection_fee_)
}

// SetUpMarketplace is a paid mutator transaction binding the contract method 0xf82ffff2.
//
// Solidity: function SetUpMarketplace(address nft_contract_, uint8 standard_, address collection_owner_, uint256 collection_fee_) returns()
func (_MetaMarketplace *MetaMarketplaceSession) SetUpMarketplace(nft_contract_ common.Address, standard_ uint8, collection_owner_ common.Address, collection_fee_ *big.Int) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.SetUpMarketplace(&_MetaMarketplace.TransactOpts, nft_contract_, standard_, collection_owner_, collection_fee_)
}

// SetUpMarketplace is a paid mutator transaction binding the contract method 0xf82ffff2.
//
// Solidity: function SetUpMarketplace(address nft_contract_, uint8 standard_, address collection_owner_, uint256 collection_fee_) returns()
func (_MetaMarketplace *MetaMarketplaceTransactorSession) SetUpMarketplace(nft_contract_ common.Address, standard_ uint8, collection_owner_ common.Address, collection_fee_ *big.Int) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.SetUpMarketplace(&_MetaMarketplace.TransactOpts, nft_contract_, standard_, collection_owner_, collection_fee_)
}

// AcceptBuyOffer is a paid mutator transaction binding the contract method 0x57194e12.
//
// Solidity: function acceptBuyOffer(address token_contract_, uint256 tokenId, uint8 currency_) returns()
func (_MetaMarketplace *MetaMarketplaceTransactor) AcceptBuyOffer(opts *bind.TransactOpts, token_contract_ common.Address, tokenId *big.Int, currency_ uint8) (*types.Transaction, error) {
	return _MetaMarketplace.contract.Transact(opts, "acceptBuyOffer", token_contract_, tokenId, currency_)
}

// AcceptBuyOffer is a paid mutator transaction binding the contract method 0x57194e12.
//
// Solidity: function acceptBuyOffer(address token_contract_, uint256 tokenId, uint8 currency_) returns()
func (_MetaMarketplace *MetaMarketplaceSession) AcceptBuyOffer(token_contract_ common.Address, tokenId *big.Int, currency_ uint8) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.AcceptBuyOffer(&_MetaMarketplace.TransactOpts, token_contract_, tokenId, currency_)
}

// AcceptBuyOffer is a paid mutator transaction binding the contract method 0x57194e12.
//
// Solidity: function acceptBuyOffer(address token_contract_, uint256 tokenId, uint8 currency_) returns()
func (_MetaMarketplace *MetaMarketplaceTransactorSession) AcceptBuyOffer(token_contract_ common.Address, tokenId *big.Int, currency_ uint8) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.AcceptBuyOffer(&_MetaMarketplace.TransactOpts, token_contract_, tokenId, currency_)
}

// EditMarketplace is a paid mutator transaction binding the contract method 0x22bb5fca.
//
// Solidity: function editMarketplace(address nft_contract_, address collection_owner_, uint256 collection_fee_) returns()
func (_MetaMarketplace *MetaMarketplaceTransactor) EditMarketplace(opts *bind.TransactOpts, nft_contract_ common.Address, collection_owner_ common.Address, collection_fee_ *big.Int) (*types.Transaction, error) {
	return _MetaMarketplace.contract.Transact(opts, "editMarketplace", nft_contract_, collection_owner_, collection_fee_)
}

// EditMarketplace is a paid mutator transaction binding the contract method 0x22bb5fca.
//
// Solidity: function editMarketplace(address nft_contract_, address collection_owner_, uint256 collection_fee_) returns()
func (_MetaMarketplace *MetaMarketplaceSession) EditMarketplace(nft_contract_ common.Address, collection_owner_ common.Address, collection_fee_ *big.Int) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.EditMarketplace(&_MetaMarketplace.TransactOpts, nft_contract_, collection_owner_, collection_fee_)
}

// EditMarketplace is a paid mutator transaction binding the contract method 0x22bb5fca.
//
// Solidity: function editMarketplace(address nft_contract_, address collection_owner_, uint256 collection_fee_) returns()
func (_MetaMarketplace *MetaMarketplaceTransactorSession) EditMarketplace(nft_contract_ common.Address, collection_owner_ common.Address, collection_fee_ *big.Int) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.EditMarketplace(&_MetaMarketplace.TransactOpts, nft_contract_, collection_owner_, collection_fee_)
}

// MakeBuyOffer is a paid mutator transaction binding the contract method 0x8813a44b.
//
// Solidity: function makeBuyOffer(address token_contract_, uint256 tokenId, uint8 currency_, uint256 bid_price_) returns()
func (_MetaMarketplace *MetaMarketplaceTransactor) MakeBuyOffer(opts *bind.TransactOpts, token_contract_ common.Address, tokenId *big.Int, currency_ uint8, bid_price_ *big.Int) (*types.Transaction, error) {
	return _MetaMarketplace.contract.Transact(opts, "makeBuyOffer", token_contract_, tokenId, currency_, bid_price_)
}

// MakeBuyOffer is a paid mutator transaction binding the contract method 0x8813a44b.
//
// Solidity: function makeBuyOffer(address token_contract_, uint256 tokenId, uint8 currency_, uint256 bid_price_) returns()
func (_MetaMarketplace *MetaMarketplaceSession) MakeBuyOffer(token_contract_ common.Address, tokenId *big.Int, currency_ uint8, bid_price_ *big.Int) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.MakeBuyOffer(&_MetaMarketplace.TransactOpts, token_contract_, tokenId, currency_, bid_price_)
}

// MakeBuyOffer is a paid mutator transaction binding the contract method 0x8813a44b.
//
// Solidity: function makeBuyOffer(address token_contract_, uint256 tokenId, uint8 currency_, uint256 bid_price_) returns()
func (_MetaMarketplace *MetaMarketplaceTransactorSession) MakeBuyOffer(token_contract_ common.Address, tokenId *big.Int, currency_ uint8, bid_price_ *big.Int) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.MakeBuyOffer(&_MetaMarketplace.TransactOpts, token_contract_, tokenId, currency_, bid_price_)
}

// MakeSellOffer is a paid mutator transaction binding the contract method 0x4ea5147e.
//
// Solidity: function makeSellOffer(uint256 tokenId, uint256 minPrice, address nft_contract_, uint8 currency_) returns()
func (_MetaMarketplace *MetaMarketplaceTransactor) MakeSellOffer(opts *bind.TransactOpts, tokenId *big.Int, minPrice *big.Int, nft_contract_ common.Address, currency_ uint8) (*types.Transaction, error) {
	return _MetaMarketplace.contract.Transact(opts, "makeSellOffer", tokenId, minPrice, nft_contract_, currency_)
}

// MakeSellOffer is a paid mutator transaction binding the contract method 0x4ea5147e.
//
// Solidity: function makeSellOffer(uint256 tokenId, uint256 minPrice, address nft_contract_, uint8 currency_) returns()
func (_MetaMarketplace *MetaMarketplaceSession) MakeSellOffer(tokenId *big.Int, minPrice *big.Int, nft_contract_ common.Address, currency_ uint8) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.MakeSellOffer(&_MetaMarketplace.TransactOpts, tokenId, minPrice, nft_contract_, currency_)
}

// MakeSellOffer is a paid mutator transaction binding the contract method 0x4ea5147e.
//
// Solidity: function makeSellOffer(uint256 tokenId, uint256 minPrice, address nft_contract_, uint8 currency_) returns()
func (_MetaMarketplace *MetaMarketplaceTransactorSession) MakeSellOffer(tokenId *big.Int, minPrice *big.Int, nft_contract_ common.Address, currency_ uint8) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.MakeSellOffer(&_MetaMarketplace.TransactOpts, tokenId, minPrice, nft_contract_, currency_)
}

// Purchase is a paid mutator transaction binding the contract method 0xfa7a5405.
//
// Solidity: function purchase(address token_contract_, uint256 tokenId, uint8 currency_, uint256 bid_price_) returns()
func (_MetaMarketplace *MetaMarketplaceTransactor) Purchase(opts *bind.TransactOpts, token_contract_ common.Address, tokenId *big.Int, currency_ uint8, bid_price_ *big.Int) (*types.Transaction, error) {
	return _MetaMarketplace.contract.Transact(opts, "purchase", token_contract_, tokenId, currency_, bid_price_)
}

// Purchase is a paid mutator transaction binding the contract method 0xfa7a5405.
//
// Solidity: function purchase(address token_contract_, uint256 tokenId, uint8 currency_, uint256 bid_price_) returns()
func (_MetaMarketplace *MetaMarketplaceSession) Purchase(token_contract_ common.Address, tokenId *big.Int, currency_ uint8, bid_price_ *big.Int) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.Purchase(&_MetaMarketplace.TransactOpts, token_contract_, tokenId, currency_, bid_price_)
}

// Purchase is a paid mutator transaction binding the contract method 0xfa7a5405.
//
// Solidity: function purchase(address token_contract_, uint256 tokenId, uint8 currency_, uint256 bid_price_) returns()
func (_MetaMarketplace *MetaMarketplaceTransactorSession) Purchase(token_contract_ common.Address, tokenId *big.Int, currency_ uint8, bid_price_ *big.Int) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.Purchase(&_MetaMarketplace.TransactOpts, token_contract_, tokenId, currency_, bid_price_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetaMarketplace *MetaMarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaMarketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetaMarketplace *MetaMarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _MetaMarketplace.Contract.RenounceOwnership(&_MetaMarketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MetaMarketplace *MetaMarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MetaMarketplace.Contract.RenounceOwnership(&_MetaMarketplace.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetaMarketplace *MetaMarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MetaMarketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetaMarketplace *MetaMarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.TransferOwnership(&_MetaMarketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MetaMarketplace *MetaMarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.TransferOwnership(&_MetaMarketplace.TransactOpts, newOwner)
}

// WithdrawBuyOffer is a paid mutator transaction binding the contract method 0x433231f3.
//
// Solidity: function withdrawBuyOffer(address token_contract_, uint256 tokenId, uint8 currency_) returns()
func (_MetaMarketplace *MetaMarketplaceTransactor) WithdrawBuyOffer(opts *bind.TransactOpts, token_contract_ common.Address, tokenId *big.Int, currency_ uint8) (*types.Transaction, error) {
	return _MetaMarketplace.contract.Transact(opts, "withdrawBuyOffer", token_contract_, tokenId, currency_)
}

// WithdrawBuyOffer is a paid mutator transaction binding the contract method 0x433231f3.
//
// Solidity: function withdrawBuyOffer(address token_contract_, uint256 tokenId, uint8 currency_) returns()
func (_MetaMarketplace *MetaMarketplaceSession) WithdrawBuyOffer(token_contract_ common.Address, tokenId *big.Int, currency_ uint8) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.WithdrawBuyOffer(&_MetaMarketplace.TransactOpts, token_contract_, tokenId, currency_)
}

// WithdrawBuyOffer is a paid mutator transaction binding the contract method 0x433231f3.
//
// Solidity: function withdrawBuyOffer(address token_contract_, uint256 tokenId, uint8 currency_) returns()
func (_MetaMarketplace *MetaMarketplaceTransactorSession) WithdrawBuyOffer(token_contract_ common.Address, tokenId *big.Int, currency_ uint8) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.WithdrawBuyOffer(&_MetaMarketplace.TransactOpts, token_contract_, tokenId, currency_)
}

// WithdrawSellOffer is a paid mutator transaction binding the contract method 0x0e079c4d.
//
// Solidity: function withdrawSellOffer(address nft_contract_, uint256 tokenId) returns()
func (_MetaMarketplace *MetaMarketplaceTransactor) WithdrawSellOffer(opts *bind.TransactOpts, nft_contract_ common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MetaMarketplace.contract.Transact(opts, "withdrawSellOffer", nft_contract_, tokenId)
}

// WithdrawSellOffer is a paid mutator transaction binding the contract method 0x0e079c4d.
//
// Solidity: function withdrawSellOffer(address nft_contract_, uint256 tokenId) returns()
func (_MetaMarketplace *MetaMarketplaceSession) WithdrawSellOffer(nft_contract_ common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.WithdrawSellOffer(&_MetaMarketplace.TransactOpts, nft_contract_, tokenId)
}

// WithdrawSellOffer is a paid mutator transaction binding the contract method 0x0e079c4d.
//
// Solidity: function withdrawSellOffer(address nft_contract_, uint256 tokenId) returns()
func (_MetaMarketplace *MetaMarketplaceTransactorSession) WithdrawSellOffer(nft_contract_ common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MetaMarketplace.Contract.WithdrawSellOffer(&_MetaMarketplace.TransactOpts, nft_contract_, tokenId)
}

// MetaMarketplaceBuyOfferWithdrawnIterator is returned from FilterBuyOfferWithdrawn and is used to iterate over the raw logs and unpacked data for BuyOfferWithdrawn events raised by the MetaMarketplace contract.
type MetaMarketplaceBuyOfferWithdrawnIterator struct {
	Event *MetaMarketplaceBuyOfferWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMarketplaceBuyOfferWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMarketplaceBuyOfferWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMarketplaceBuyOfferWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMarketplaceBuyOfferWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMarketplaceBuyOfferWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMarketplaceBuyOfferWithdrawn represents a BuyOfferWithdrawn event raised by the MetaMarketplace contract.
type MetaMarketplaceBuyOfferWithdrawn struct {
	TokenId *big.Int
	Buyer   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBuyOfferWithdrawn is a free log retrieval operation binding the contract event 0xcde8efb885f61e3023afcfb4e801eca4790b08b4d7706ea9e7fb62073a09574a.
//
// Solidity: event BuyOfferWithdrawn(uint256 tokenId, address buyer)
func (_MetaMarketplace *MetaMarketplaceFilterer) FilterBuyOfferWithdrawn(opts *bind.FilterOpts) (*MetaMarketplaceBuyOfferWithdrawnIterator, error) {

	logs, sub, err := _MetaMarketplace.contract.FilterLogs(opts, "BuyOfferWithdrawn")
	if err != nil {
		return nil, err
	}
	return &MetaMarketplaceBuyOfferWithdrawnIterator{contract: _MetaMarketplace.contract, event: "BuyOfferWithdrawn", logs: logs, sub: sub}, nil
}

// WatchBuyOfferWithdrawn is a free log subscription operation binding the contract event 0xcde8efb885f61e3023afcfb4e801eca4790b08b4d7706ea9e7fb62073a09574a.
//
// Solidity: event BuyOfferWithdrawn(uint256 tokenId, address buyer)
func (_MetaMarketplace *MetaMarketplaceFilterer) WatchBuyOfferWithdrawn(opts *bind.WatchOpts, sink chan<- *MetaMarketplaceBuyOfferWithdrawn) (event.Subscription, error) {

	logs, sub, err := _MetaMarketplace.contract.WatchLogs(opts, "BuyOfferWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMarketplaceBuyOfferWithdrawn)
				if err := _MetaMarketplace.contract.UnpackLog(event, "BuyOfferWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBuyOfferWithdrawn is a log parse operation binding the contract event 0xcde8efb885f61e3023afcfb4e801eca4790b08b4d7706ea9e7fb62073a09574a.
//
// Solidity: event BuyOfferWithdrawn(uint256 tokenId, address buyer)
func (_MetaMarketplace *MetaMarketplaceFilterer) ParseBuyOfferWithdrawn(log types.Log) (*MetaMarketplaceBuyOfferWithdrawn, error) {
	event := new(MetaMarketplaceBuyOfferWithdrawn)
	if err := _MetaMarketplace.contract.UnpackLog(event, "BuyOfferWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMarketplaceCalculatedFeesIterator is returned from FilterCalculatedFees and is used to iterate over the raw logs and unpacked data for CalculatedFees events raised by the MetaMarketplace contract.
type MetaMarketplaceCalculatedFeesIterator struct {
	Event *MetaMarketplaceCalculatedFees // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMarketplaceCalculatedFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMarketplaceCalculatedFees)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMarketplaceCalculatedFees)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMarketplaceCalculatedFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMarketplaceCalculatedFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMarketplaceCalculatedFees represents a CalculatedFees event raised by the MetaMarketplace contract.
type MetaMarketplaceCalculatedFees struct {
	InitialValue     *big.Int
	Fees             *big.Int
	TransferedAmount *big.Int
	FeeAddress       common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCalculatedFees is a free log retrieval operation binding the contract event 0xe1ac6cda82b822e7f388493f39f37748308acbc0a4c3044861e00a2ffaa46f9e.
//
// Solidity: event CalculatedFees(uint256 initial_value, uint256 fees, uint256 transfered_amount, address feeAddress)
func (_MetaMarketplace *MetaMarketplaceFilterer) FilterCalculatedFees(opts *bind.FilterOpts) (*MetaMarketplaceCalculatedFeesIterator, error) {

	logs, sub, err := _MetaMarketplace.contract.FilterLogs(opts, "CalculatedFees")
	if err != nil {
		return nil, err
	}
	return &MetaMarketplaceCalculatedFeesIterator{contract: _MetaMarketplace.contract, event: "CalculatedFees", logs: logs, sub: sub}, nil
}

// WatchCalculatedFees is a free log subscription operation binding the contract event 0xe1ac6cda82b822e7f388493f39f37748308acbc0a4c3044861e00a2ffaa46f9e.
//
// Solidity: event CalculatedFees(uint256 initial_value, uint256 fees, uint256 transfered_amount, address feeAddress)
func (_MetaMarketplace *MetaMarketplaceFilterer) WatchCalculatedFees(opts *bind.WatchOpts, sink chan<- *MetaMarketplaceCalculatedFees) (event.Subscription, error) {

	logs, sub, err := _MetaMarketplace.contract.WatchLogs(opts, "CalculatedFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMarketplaceCalculatedFees)
				if err := _MetaMarketplace.contract.UnpackLog(event, "CalculatedFees", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCalculatedFees is a log parse operation binding the contract event 0xe1ac6cda82b822e7f388493f39f37748308acbc0a4c3044861e00a2ffaa46f9e.
//
// Solidity: event CalculatedFees(uint256 initial_value, uint256 fees, uint256 transfered_amount, address feeAddress)
func (_MetaMarketplace *MetaMarketplaceFilterer) ParseCalculatedFees(log types.Log) (*MetaMarketplaceCalculatedFees, error) {
	event := new(MetaMarketplaceCalculatedFees)
	if err := _MetaMarketplace.contract.UnpackLog(event, "CalculatedFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMarketplaceNewBuyOfferIterator is returned from FilterNewBuyOffer and is used to iterate over the raw logs and unpacked data for NewBuyOffer events raised by the MetaMarketplace contract.
type MetaMarketplaceNewBuyOfferIterator struct {
	Event *MetaMarketplaceNewBuyOffer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMarketplaceNewBuyOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMarketplaceNewBuyOffer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMarketplaceNewBuyOffer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMarketplaceNewBuyOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMarketplaceNewBuyOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMarketplaceNewBuyOffer represents a NewBuyOffer event raised by the MetaMarketplace contract.
type MetaMarketplaceNewBuyOffer struct {
	TokenId *big.Int
	Buyer   common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewBuyOffer is a free log retrieval operation binding the contract event 0xd1d0f457f709cbed0f3360dee6050dde1a0ac4706268ff21aecae87e6ddb2861.
//
// Solidity: event NewBuyOffer(uint256 tokenId, address buyer, uint256 value)
func (_MetaMarketplace *MetaMarketplaceFilterer) FilterNewBuyOffer(opts *bind.FilterOpts) (*MetaMarketplaceNewBuyOfferIterator, error) {

	logs, sub, err := _MetaMarketplace.contract.FilterLogs(opts, "NewBuyOffer")
	if err != nil {
		return nil, err
	}
	return &MetaMarketplaceNewBuyOfferIterator{contract: _MetaMarketplace.contract, event: "NewBuyOffer", logs: logs, sub: sub}, nil
}

// WatchNewBuyOffer is a free log subscription operation binding the contract event 0xd1d0f457f709cbed0f3360dee6050dde1a0ac4706268ff21aecae87e6ddb2861.
//
// Solidity: event NewBuyOffer(uint256 tokenId, address buyer, uint256 value)
func (_MetaMarketplace *MetaMarketplaceFilterer) WatchNewBuyOffer(opts *bind.WatchOpts, sink chan<- *MetaMarketplaceNewBuyOffer) (event.Subscription, error) {

	logs, sub, err := _MetaMarketplace.contract.WatchLogs(opts, "NewBuyOffer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMarketplaceNewBuyOffer)
				if err := _MetaMarketplace.contract.UnpackLog(event, "NewBuyOffer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBuyOffer is a log parse operation binding the contract event 0xd1d0f457f709cbed0f3360dee6050dde1a0ac4706268ff21aecae87e6ddb2861.
//
// Solidity: event NewBuyOffer(uint256 tokenId, address buyer, uint256 value)
func (_MetaMarketplace *MetaMarketplaceFilterer) ParseNewBuyOffer(log types.Log) (*MetaMarketplaceNewBuyOffer, error) {
	event := new(MetaMarketplaceNewBuyOffer)
	if err := _MetaMarketplace.contract.UnpackLog(event, "NewBuyOffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMarketplaceNewSellOfferIterator is returned from FilterNewSellOffer and is used to iterate over the raw logs and unpacked data for NewSellOffer events raised by the MetaMarketplace contract.
type MetaMarketplaceNewSellOfferIterator struct {
	Event *MetaMarketplaceNewSellOffer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMarketplaceNewSellOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMarketplaceNewSellOffer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMarketplaceNewSellOffer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMarketplaceNewSellOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMarketplaceNewSellOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMarketplaceNewSellOffer represents a NewSellOffer event raised by the MetaMarketplace contract.
type MetaMarketplaceNewSellOffer struct {
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewSellOffer is a free log retrieval operation binding the contract event 0xff04060fdb1ade3f35e2fbbdae70256ab6294fe60ce3e63dfcd01d39a73e462f.
//
// Solidity: event NewSellOffer(address nft_contract_, uint256 tokenId, address seller, uint256 value)
func (_MetaMarketplace *MetaMarketplaceFilterer) FilterNewSellOffer(opts *bind.FilterOpts) (*MetaMarketplaceNewSellOfferIterator, error) {

	logs, sub, err := _MetaMarketplace.contract.FilterLogs(opts, "NewSellOffer")
	if err != nil {
		return nil, err
	}
	return &MetaMarketplaceNewSellOfferIterator{contract: _MetaMarketplace.contract, event: "NewSellOffer", logs: logs, sub: sub}, nil
}

// WatchNewSellOffer is a free log subscription operation binding the contract event 0xff04060fdb1ade3f35e2fbbdae70256ab6294fe60ce3e63dfcd01d39a73e462f.
//
// Solidity: event NewSellOffer(address nft_contract_, uint256 tokenId, address seller, uint256 value)
func (_MetaMarketplace *MetaMarketplaceFilterer) WatchNewSellOffer(opts *bind.WatchOpts, sink chan<- *MetaMarketplaceNewSellOffer) (event.Subscription, error) {

	logs, sub, err := _MetaMarketplace.contract.WatchLogs(opts, "NewSellOffer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMarketplaceNewSellOffer)
				if err := _MetaMarketplace.contract.UnpackLog(event, "NewSellOffer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewSellOffer is a log parse operation binding the contract event 0xff04060fdb1ade3f35e2fbbdae70256ab6294fe60ce3e63dfcd01d39a73e462f.
//
// Solidity: event NewSellOffer(address nft_contract_, uint256 tokenId, address seller, uint256 value)
func (_MetaMarketplace *MetaMarketplaceFilterer) ParseNewSellOffer(log types.Log) (*MetaMarketplaceNewSellOffer, error) {
	event := new(MetaMarketplaceNewSellOffer)
	if err := _MetaMarketplace.contract.UnpackLog(event, "NewSellOffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MetaMarketplace contract.
type MetaMarketplaceOwnershipTransferredIterator struct {
	Event *MetaMarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMarketplaceOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMarketplaceOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the MetaMarketplace contract.
type MetaMarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetaMarketplace *MetaMarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MetaMarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MetaMarketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MetaMarketplaceOwnershipTransferredIterator{contract: _MetaMarketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetaMarketplace *MetaMarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MetaMarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MetaMarketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMarketplaceOwnershipTransferred)
				if err := _MetaMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MetaMarketplace *MetaMarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*MetaMarketplaceOwnershipTransferred, error) {
	event := new(MetaMarketplaceOwnershipTransferred)
	if err := _MetaMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMarketplaceRoyaltiesPaidIterator is returned from FilterRoyaltiesPaid and is used to iterate over the raw logs and unpacked data for RoyaltiesPaid events raised by the MetaMarketplace contract.
type MetaMarketplaceRoyaltiesPaidIterator struct {
	Event *MetaMarketplaceRoyaltiesPaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMarketplaceRoyaltiesPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMarketplaceRoyaltiesPaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMarketplaceRoyaltiesPaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMarketplaceRoyaltiesPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMarketplaceRoyaltiesPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMarketplaceRoyaltiesPaid represents a RoyaltiesPaid event raised by the MetaMarketplace contract.
type MetaMarketplaceRoyaltiesPaid struct {
	NftContract common.Address
	TokenId     *big.Int
	Recepient   common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRoyaltiesPaid is a free log retrieval operation binding the contract event 0x5825d7670ca9035e2d333fd7f1cfca7d8e476c3522beb6447f33885ca68a5ed1.
//
// Solidity: event RoyaltiesPaid(address nft_contract_, uint256 tokenId, address recepient, uint256 value)
func (_MetaMarketplace *MetaMarketplaceFilterer) FilterRoyaltiesPaid(opts *bind.FilterOpts) (*MetaMarketplaceRoyaltiesPaidIterator, error) {

	logs, sub, err := _MetaMarketplace.contract.FilterLogs(opts, "RoyaltiesPaid")
	if err != nil {
		return nil, err
	}
	return &MetaMarketplaceRoyaltiesPaidIterator{contract: _MetaMarketplace.contract, event: "RoyaltiesPaid", logs: logs, sub: sub}, nil
}

// WatchRoyaltiesPaid is a free log subscription operation binding the contract event 0x5825d7670ca9035e2d333fd7f1cfca7d8e476c3522beb6447f33885ca68a5ed1.
//
// Solidity: event RoyaltiesPaid(address nft_contract_, uint256 tokenId, address recepient, uint256 value)
func (_MetaMarketplace *MetaMarketplaceFilterer) WatchRoyaltiesPaid(opts *bind.WatchOpts, sink chan<- *MetaMarketplaceRoyaltiesPaid) (event.Subscription, error) {

	logs, sub, err := _MetaMarketplace.contract.WatchLogs(opts, "RoyaltiesPaid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMarketplaceRoyaltiesPaid)
				if err := _MetaMarketplace.contract.UnpackLog(event, "RoyaltiesPaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoyaltiesPaid is a log parse operation binding the contract event 0x5825d7670ca9035e2d333fd7f1cfca7d8e476c3522beb6447f33885ca68a5ed1.
//
// Solidity: event RoyaltiesPaid(address nft_contract_, uint256 tokenId, address recepient, uint256 value)
func (_MetaMarketplace *MetaMarketplaceFilterer) ParseRoyaltiesPaid(log types.Log) (*MetaMarketplaceRoyaltiesPaid, error) {
	event := new(MetaMarketplaceRoyaltiesPaid)
	if err := _MetaMarketplace.contract.UnpackLog(event, "RoyaltiesPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMarketplaceSaleIterator is returned from FilterSale and is used to iterate over the raw logs and unpacked data for Sale events raised by the MetaMarketplace contract.
type MetaMarketplaceSaleIterator struct {
	Event *MetaMarketplaceSale // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMarketplaceSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMarketplaceSale)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMarketplaceSale)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMarketplaceSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMarketplaceSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMarketplaceSale represents a Sale event raised by the MetaMarketplace contract.
type MetaMarketplaceSale struct {
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Buyer       common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSale is a free log retrieval operation binding the contract event 0x3644379fe03acf3b2db463b2b1fb79bfa0a162795d0e7bb0c8a999bd98db8135.
//
// Solidity: event Sale(address nft_contract_, uint256 tokenId, address seller, address buyer, uint256 value)
func (_MetaMarketplace *MetaMarketplaceFilterer) FilterSale(opts *bind.FilterOpts) (*MetaMarketplaceSaleIterator, error) {

	logs, sub, err := _MetaMarketplace.contract.FilterLogs(opts, "Sale")
	if err != nil {
		return nil, err
	}
	return &MetaMarketplaceSaleIterator{contract: _MetaMarketplace.contract, event: "Sale", logs: logs, sub: sub}, nil
}

// WatchSale is a free log subscription operation binding the contract event 0x3644379fe03acf3b2db463b2b1fb79bfa0a162795d0e7bb0c8a999bd98db8135.
//
// Solidity: event Sale(address nft_contract_, uint256 tokenId, address seller, address buyer, uint256 value)
func (_MetaMarketplace *MetaMarketplaceFilterer) WatchSale(opts *bind.WatchOpts, sink chan<- *MetaMarketplaceSale) (event.Subscription, error) {

	logs, sub, err := _MetaMarketplace.contract.WatchLogs(opts, "Sale")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMarketplaceSale)
				if err := _MetaMarketplace.contract.UnpackLog(event, "Sale", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSale is a log parse operation binding the contract event 0x3644379fe03acf3b2db463b2b1fb79bfa0a162795d0e7bb0c8a999bd98db8135.
//
// Solidity: event Sale(address nft_contract_, uint256 tokenId, address seller, address buyer, uint256 value)
func (_MetaMarketplace *MetaMarketplaceFilterer) ParseSale(log types.Log) (*MetaMarketplaceSale, error) {
	event := new(MetaMarketplaceSale)
	if err := _MetaMarketplace.contract.UnpackLog(event, "Sale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaMarketplaceSellOfferWithdrawnIterator is returned from FilterSellOfferWithdrawn and is used to iterate over the raw logs and unpacked data for SellOfferWithdrawn events raised by the MetaMarketplace contract.
type MetaMarketplaceSellOfferWithdrawnIterator struct {
	Event *MetaMarketplaceSellOfferWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaMarketplaceSellOfferWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaMarketplaceSellOfferWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaMarketplaceSellOfferWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaMarketplaceSellOfferWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaMarketplaceSellOfferWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaMarketplaceSellOfferWithdrawn represents a SellOfferWithdrawn event raised by the MetaMarketplace contract.
type MetaMarketplaceSellOfferWithdrawn struct {
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSellOfferWithdrawn is a free log retrieval operation binding the contract event 0xce00375f5a3b9c4a857fc9fa1b243f1af9b10a7db9d2d47227fe8217003c9ad4.
//
// Solidity: event SellOfferWithdrawn(address nft_contract_, uint256 tokenId, address seller)
func (_MetaMarketplace *MetaMarketplaceFilterer) FilterSellOfferWithdrawn(opts *bind.FilterOpts) (*MetaMarketplaceSellOfferWithdrawnIterator, error) {

	logs, sub, err := _MetaMarketplace.contract.FilterLogs(opts, "SellOfferWithdrawn")
	if err != nil {
		return nil, err
	}
	return &MetaMarketplaceSellOfferWithdrawnIterator{contract: _MetaMarketplace.contract, event: "SellOfferWithdrawn", logs: logs, sub: sub}, nil
}

// WatchSellOfferWithdrawn is a free log subscription operation binding the contract event 0xce00375f5a3b9c4a857fc9fa1b243f1af9b10a7db9d2d47227fe8217003c9ad4.
//
// Solidity: event SellOfferWithdrawn(address nft_contract_, uint256 tokenId, address seller)
func (_MetaMarketplace *MetaMarketplaceFilterer) WatchSellOfferWithdrawn(opts *bind.WatchOpts, sink chan<- *MetaMarketplaceSellOfferWithdrawn) (event.Subscription, error) {

	logs, sub, err := _MetaMarketplace.contract.WatchLogs(opts, "SellOfferWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaMarketplaceSellOfferWithdrawn)
				if err := _MetaMarketplace.contract.UnpackLog(event, "SellOfferWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSellOfferWithdrawn is a log parse operation binding the contract event 0xce00375f5a3b9c4a857fc9fa1b243f1af9b10a7db9d2d47227fe8217003c9ad4.
//
// Solidity: event SellOfferWithdrawn(address nft_contract_, uint256 tokenId, address seller)
func (_MetaMarketplace *MetaMarketplaceFilterer) ParseSellOfferWithdrawn(log types.Log) (*MetaMarketplaceSellOfferWithdrawn, error) {
	event := new(MetaMarketplaceSellOfferWithdrawn)
	if err := _MetaMarketplace.contract.UnpackLog(event, "SellOfferWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
