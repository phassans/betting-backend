// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package betting

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
	_ = abi.ConvertType
)

// BettingBetInfo is an auto generated low-level Go binding around an user-defined struct.
type BettingBetInfo struct {
	Id           *big.Int
	UserA        common.Address
	UserB        common.Address
	Amount       *big.Int
	Winner       common.Address
	Reward       *big.Int
	IsLong       bool
	CreateTime   *big.Int
	ExpireTime   *big.Int
	ClosingTime  *big.Int
	OpeningPrice *big.Int
	ClosingPrice *big.Int
	BetStatus    *big.Int
}

// BettingMetaData contains all meta data concerning the Betting contract.
var BettingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_botAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_botFeeBasisPoints\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userB\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"openingPrice\",\"type\":\"uint256\"}],\"name\":\"BetActived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"closingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"BetClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"betId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userA\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BetRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userA\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expireTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"closingTime\",\"type\":\"uint256\"}],\"name\":\"NewBetCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardWithdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isAllowed\",\"type\":\"bool\"}],\"name\":\"allowUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"betCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_usdcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expireTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_closingTime\",\"type\":\"uint256\"}],\"name\":\"createBet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveBetIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBTCPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"}],\"name\":\"getBetIdsForUser\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_betID\",\"type\":\"uint256\"}],\"name\":\"getBetInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"userA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"userB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betStatus\",\"type\":\"uint256\"}],\"internalType\":\"structBetting.BetInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBotAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBotFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingBetIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceFeedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUSDCAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_betID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_usdcAmount\",\"type\":\"uint256\"}],\"name\":\"joinBet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"betId\",\"type\":\"uint256\"}],\"name\":\"refundBet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_betID\",\"type\":\"uint256\"}],\"name\":\"resolveBet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_botAddress\",\"type\":\"address\"}],\"name\":\"setBotAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_botFeeBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setBotFeeBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feedAddress\",\"type\":\"address\"}],\"name\":\"setPriceFeedAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdcAddress\",\"type\":\"address\"}],\"name\":\"setUSDCAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_betID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BettingABI is the input ABI used to generate the binding from.
// Deprecated: Use BettingMetaData.ABI instead.
var BettingABI = BettingMetaData.ABI

// Betting is an auto generated Go binding around an Ethereum contract.
type Betting struct {
	BettingCaller     // Read-only binding to the contract
	BettingTransactor // Write-only binding to the contract
	BettingFilterer   // Log filterer for contract events
}

// BettingCaller is an auto generated read-only Go binding around an Ethereum contract.
type BettingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BettingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BettingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BettingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BettingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BettingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BettingSession struct {
	Contract     *Betting          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BettingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BettingCallerSession struct {
	Contract *BettingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BettingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BettingTransactorSession struct {
	Contract     *BettingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BettingRaw is an auto generated low-level Go binding around an Ethereum contract.
type BettingRaw struct {
	Contract *Betting // Generic contract binding to access the raw methods on
}

// BettingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BettingCallerRaw struct {
	Contract *BettingCaller // Generic read-only contract binding to access the raw methods on
}

// BettingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BettingTransactorRaw struct {
	Contract *BettingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBetting creates a new instance of Betting, bound to a specific deployed contract.
func NewBetting(address common.Address, backend bind.ContractBackend) (*Betting, error) {
	contract, err := bindBetting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Betting{BettingCaller: BettingCaller{contract: contract}, BettingTransactor: BettingTransactor{contract: contract}, BettingFilterer: BettingFilterer{contract: contract}}, nil
}

// NewBettingCaller creates a new read-only instance of Betting, bound to a specific deployed contract.
func NewBettingCaller(address common.Address, caller bind.ContractCaller) (*BettingCaller, error) {
	contract, err := bindBetting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BettingCaller{contract: contract}, nil
}

// NewBettingTransactor creates a new write-only instance of Betting, bound to a specific deployed contract.
func NewBettingTransactor(address common.Address, transactor bind.ContractTransactor) (*BettingTransactor, error) {
	contract, err := bindBetting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BettingTransactor{contract: contract}, nil
}

// NewBettingFilterer creates a new log filterer instance of Betting, bound to a specific deployed contract.
func NewBettingFilterer(address common.Address, filterer bind.ContractFilterer) (*BettingFilterer, error) {
	contract, err := bindBetting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BettingFilterer{contract: contract}, nil
}

// bindBetting binds a generic wrapper to an already deployed contract.
func bindBetting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BettingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Betting *BettingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Betting.Contract.BettingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Betting *BettingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting.Contract.BettingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Betting *BettingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Betting.Contract.BettingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Betting *BettingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Betting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Betting *BettingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Betting *BettingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Betting.Contract.contract.Transact(opts, method, params...)
}

// BetCount is a free data retrieval call binding the contract method 0xcf094497.
//
// Solidity: function betCount() view returns(uint256)
func (_Betting *BettingCaller) BetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "betCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BetCount is a free data retrieval call binding the contract method 0xcf094497.
//
// Solidity: function betCount() view returns(uint256)
func (_Betting *BettingSession) BetCount() (*big.Int, error) {
	return _Betting.Contract.BetCount(&_Betting.CallOpts)
}

// BetCount is a free data retrieval call binding the contract method 0xcf094497.
//
// Solidity: function betCount() view returns(uint256)
func (_Betting *BettingCallerSession) BetCount() (*big.Int, error) {
	return _Betting.Contract.BetCount(&_Betting.CallOpts)
}

// GetActiveBetIDs is a free data retrieval call binding the contract method 0xe2587a11.
//
// Solidity: function getActiveBetIDs() view returns(uint256[])
func (_Betting *BettingCaller) GetActiveBetIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "getActiveBetIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetActiveBetIDs is a free data retrieval call binding the contract method 0xe2587a11.
//
// Solidity: function getActiveBetIDs() view returns(uint256[])
func (_Betting *BettingSession) GetActiveBetIDs() ([]*big.Int, error) {
	return _Betting.Contract.GetActiveBetIDs(&_Betting.CallOpts)
}

// GetActiveBetIDs is a free data retrieval call binding the contract method 0xe2587a11.
//
// Solidity: function getActiveBetIDs() view returns(uint256[])
func (_Betting *BettingCallerSession) GetActiveBetIDs() ([]*big.Int, error) {
	return _Betting.Contract.GetActiveBetIDs(&_Betting.CallOpts)
}

// GetBTCPrice is a free data retrieval call binding the contract method 0xdd22db80.
//
// Solidity: function getBTCPrice() view returns(uint256)
func (_Betting *BettingCaller) GetBTCPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "getBTCPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBTCPrice is a free data retrieval call binding the contract method 0xdd22db80.
//
// Solidity: function getBTCPrice() view returns(uint256)
func (_Betting *BettingSession) GetBTCPrice() (*big.Int, error) {
	return _Betting.Contract.GetBTCPrice(&_Betting.CallOpts)
}

// GetBTCPrice is a free data retrieval call binding the contract method 0xdd22db80.
//
// Solidity: function getBTCPrice() view returns(uint256)
func (_Betting *BettingCallerSession) GetBTCPrice() (*big.Int, error) {
	return _Betting.Contract.GetBTCPrice(&_Betting.CallOpts)
}

// GetBetIdsForUser is a free data retrieval call binding the contract method 0x76dd3fb1.
//
// Solidity: function getBetIdsForUser(address _userAddress) view returns(uint256[])
func (_Betting *BettingCaller) GetBetIdsForUser(opts *bind.CallOpts, _userAddress common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "getBetIdsForUser", _userAddress)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBetIdsForUser is a free data retrieval call binding the contract method 0x76dd3fb1.
//
// Solidity: function getBetIdsForUser(address _userAddress) view returns(uint256[])
func (_Betting *BettingSession) GetBetIdsForUser(_userAddress common.Address) ([]*big.Int, error) {
	return _Betting.Contract.GetBetIdsForUser(&_Betting.CallOpts, _userAddress)
}

// GetBetIdsForUser is a free data retrieval call binding the contract method 0x76dd3fb1.
//
// Solidity: function getBetIdsForUser(address _userAddress) view returns(uint256[])
func (_Betting *BettingCallerSession) GetBetIdsForUser(_userAddress common.Address) ([]*big.Int, error) {
	return _Betting.Contract.GetBetIdsForUser(&_Betting.CallOpts, _userAddress)
}

// GetBetInfo is a free data retrieval call binding the contract method 0x79141f80.
//
// Solidity: function getBetInfo(uint256 _betID) view returns((uint256,address,address,uint256,address,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Betting *BettingCaller) GetBetInfo(opts *bind.CallOpts, _betID *big.Int) (BettingBetInfo, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "getBetInfo", _betID)

	if err != nil {
		return *new(BettingBetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(BettingBetInfo)).(*BettingBetInfo)

	return out0, err

}

// GetBetInfo is a free data retrieval call binding the contract method 0x79141f80.
//
// Solidity: function getBetInfo(uint256 _betID) view returns((uint256,address,address,uint256,address,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Betting *BettingSession) GetBetInfo(_betID *big.Int) (BettingBetInfo, error) {
	return _Betting.Contract.GetBetInfo(&_Betting.CallOpts, _betID)
}

// GetBetInfo is a free data retrieval call binding the contract method 0x79141f80.
//
// Solidity: function getBetInfo(uint256 _betID) view returns((uint256,address,address,uint256,address,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Betting *BettingCallerSession) GetBetInfo(_betID *big.Int) (BettingBetInfo, error) {
	return _Betting.Contract.GetBetInfo(&_Betting.CallOpts, _betID)
}

// GetBotAddress is a free data retrieval call binding the contract method 0xc2fdda7d.
//
// Solidity: function getBotAddress() view returns(address)
func (_Betting *BettingCaller) GetBotAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "getBotAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBotAddress is a free data retrieval call binding the contract method 0xc2fdda7d.
//
// Solidity: function getBotAddress() view returns(address)
func (_Betting *BettingSession) GetBotAddress() (common.Address, error) {
	return _Betting.Contract.GetBotAddress(&_Betting.CallOpts)
}

// GetBotAddress is a free data retrieval call binding the contract method 0xc2fdda7d.
//
// Solidity: function getBotAddress() view returns(address)
func (_Betting *BettingCallerSession) GetBotAddress() (common.Address, error) {
	return _Betting.Contract.GetBotAddress(&_Betting.CallOpts)
}

// GetBotFeeBasisPoints is a free data retrieval call binding the contract method 0x3fcd83ad.
//
// Solidity: function getBotFeeBasisPoints() view returns(uint256)
func (_Betting *BettingCaller) GetBotFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "getBotFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBotFeeBasisPoints is a free data retrieval call binding the contract method 0x3fcd83ad.
//
// Solidity: function getBotFeeBasisPoints() view returns(uint256)
func (_Betting *BettingSession) GetBotFeeBasisPoints() (*big.Int, error) {
	return _Betting.Contract.GetBotFeeBasisPoints(&_Betting.CallOpts)
}

// GetBotFeeBasisPoints is a free data retrieval call binding the contract method 0x3fcd83ad.
//
// Solidity: function getBotFeeBasisPoints() view returns(uint256)
func (_Betting *BettingCallerSession) GetBotFeeBasisPoints() (*big.Int, error) {
	return _Betting.Contract.GetBotFeeBasisPoints(&_Betting.CallOpts)
}

// GetPendingBetIDs is a free data retrieval call binding the contract method 0x560410c7.
//
// Solidity: function getPendingBetIDs() view returns(uint256[])
func (_Betting *BettingCaller) GetPendingBetIDs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "getPendingBetIDs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPendingBetIDs is a free data retrieval call binding the contract method 0x560410c7.
//
// Solidity: function getPendingBetIDs() view returns(uint256[])
func (_Betting *BettingSession) GetPendingBetIDs() ([]*big.Int, error) {
	return _Betting.Contract.GetPendingBetIDs(&_Betting.CallOpts)
}

// GetPendingBetIDs is a free data retrieval call binding the contract method 0x560410c7.
//
// Solidity: function getPendingBetIDs() view returns(uint256[])
func (_Betting *BettingCallerSession) GetPendingBetIDs() ([]*big.Int, error) {
	return _Betting.Contract.GetPendingBetIDs(&_Betting.CallOpts)
}

// GetPriceFeedAddress is a free data retrieval call binding the contract method 0x0f0f30b2.
//
// Solidity: function getPriceFeedAddress() view returns(address)
func (_Betting *BettingCaller) GetPriceFeedAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "getPriceFeedAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceFeedAddress is a free data retrieval call binding the contract method 0x0f0f30b2.
//
// Solidity: function getPriceFeedAddress() view returns(address)
func (_Betting *BettingSession) GetPriceFeedAddress() (common.Address, error) {
	return _Betting.Contract.GetPriceFeedAddress(&_Betting.CallOpts)
}

// GetPriceFeedAddress is a free data retrieval call binding the contract method 0x0f0f30b2.
//
// Solidity: function getPriceFeedAddress() view returns(address)
func (_Betting *BettingCallerSession) GetPriceFeedAddress() (common.Address, error) {
	return _Betting.Contract.GetPriceFeedAddress(&_Betting.CallOpts)
}

// GetUSDCAddress is a free data retrieval call binding the contract method 0xbc06e81d.
//
// Solidity: function getUSDCAddress() view returns(address)
func (_Betting *BettingCaller) GetUSDCAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "getUSDCAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUSDCAddress is a free data retrieval call binding the contract method 0xbc06e81d.
//
// Solidity: function getUSDCAddress() view returns(address)
func (_Betting *BettingSession) GetUSDCAddress() (common.Address, error) {
	return _Betting.Contract.GetUSDCAddress(&_Betting.CallOpts)
}

// GetUSDCAddress is a free data retrieval call binding the contract method 0xbc06e81d.
//
// Solidity: function getUSDCAddress() view returns(address)
func (_Betting *BettingCallerSession) GetUSDCAddress() (common.Address, error) {
	return _Betting.Contract.GetUSDCAddress(&_Betting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Betting *BettingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Betting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Betting *BettingSession) Owner() (common.Address, error) {
	return _Betting.Contract.Owner(&_Betting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Betting *BettingCallerSession) Owner() (common.Address, error) {
	return _Betting.Contract.Owner(&_Betting.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Betting *BettingTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Betting *BettingSession) AcceptOwnership() (*types.Transaction, error) {
	return _Betting.Contract.AcceptOwnership(&_Betting.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Betting *BettingTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Betting.Contract.AcceptOwnership(&_Betting.TransactOpts)
}

// AllowUser is a paid mutator transaction binding the contract method 0x726b19c3.
//
// Solidity: function allowUser(address _userAddress, bool _isAllowed) returns()
func (_Betting *BettingTransactor) AllowUser(opts *bind.TransactOpts, _userAddress common.Address, _isAllowed bool) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "allowUser", _userAddress, _isAllowed)
}

// AllowUser is a paid mutator transaction binding the contract method 0x726b19c3.
//
// Solidity: function allowUser(address _userAddress, bool _isAllowed) returns()
func (_Betting *BettingSession) AllowUser(_userAddress common.Address, _isAllowed bool) (*types.Transaction, error) {
	return _Betting.Contract.AllowUser(&_Betting.TransactOpts, _userAddress, _isAllowed)
}

// AllowUser is a paid mutator transaction binding the contract method 0x726b19c3.
//
// Solidity: function allowUser(address _userAddress, bool _isAllowed) returns()
func (_Betting *BettingTransactorSession) AllowUser(_userAddress common.Address, _isAllowed bool) (*types.Transaction, error) {
	return _Betting.Contract.AllowUser(&_Betting.TransactOpts, _userAddress, _isAllowed)
}

// CreateBet is a paid mutator transaction binding the contract method 0x3e6c8c5f.
//
// Solidity: function createBet(bool isLong, uint256 _usdcAmount, uint256 _expireTime, uint256 _closingTime) payable returns(uint256)
func (_Betting *BettingTransactor) CreateBet(opts *bind.TransactOpts, isLong bool, _usdcAmount *big.Int, _expireTime *big.Int, _closingTime *big.Int) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "createBet", isLong, _usdcAmount, _expireTime, _closingTime)
}

// CreateBet is a paid mutator transaction binding the contract method 0x3e6c8c5f.
//
// Solidity: function createBet(bool isLong, uint256 _usdcAmount, uint256 _expireTime, uint256 _closingTime) payable returns(uint256)
func (_Betting *BettingSession) CreateBet(isLong bool, _usdcAmount *big.Int, _expireTime *big.Int, _closingTime *big.Int) (*types.Transaction, error) {
	return _Betting.Contract.CreateBet(&_Betting.TransactOpts, isLong, _usdcAmount, _expireTime, _closingTime)
}

// CreateBet is a paid mutator transaction binding the contract method 0x3e6c8c5f.
//
// Solidity: function createBet(bool isLong, uint256 _usdcAmount, uint256 _expireTime, uint256 _closingTime) payable returns(uint256)
func (_Betting *BettingTransactorSession) CreateBet(isLong bool, _usdcAmount *big.Int, _expireTime *big.Int, _closingTime *big.Int) (*types.Transaction, error) {
	return _Betting.Contract.CreateBet(&_Betting.TransactOpts, isLong, _usdcAmount, _expireTime, _closingTime)
}

// JoinBet is a paid mutator transaction binding the contract method 0x59b0a8ed.
//
// Solidity: function joinBet(uint256 _betID, uint256 _usdcAmount) payable returns(uint256)
func (_Betting *BettingTransactor) JoinBet(opts *bind.TransactOpts, _betID *big.Int, _usdcAmount *big.Int) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "joinBet", _betID, _usdcAmount)
}

// JoinBet is a paid mutator transaction binding the contract method 0x59b0a8ed.
//
// Solidity: function joinBet(uint256 _betID, uint256 _usdcAmount) payable returns(uint256)
func (_Betting *BettingSession) JoinBet(_betID *big.Int, _usdcAmount *big.Int) (*types.Transaction, error) {
	return _Betting.Contract.JoinBet(&_Betting.TransactOpts, _betID, _usdcAmount)
}

// JoinBet is a paid mutator transaction binding the contract method 0x59b0a8ed.
//
// Solidity: function joinBet(uint256 _betID, uint256 _usdcAmount) payable returns(uint256)
func (_Betting *BettingTransactorSession) JoinBet(_betID *big.Int, _usdcAmount *big.Int) (*types.Transaction, error) {
	return _Betting.Contract.JoinBet(&_Betting.TransactOpts, _betID, _usdcAmount)
}

// RefundBet is a paid mutator transaction binding the contract method 0xe1fdb4b4.
//
// Solidity: function refundBet(uint256 betId) returns()
func (_Betting *BettingTransactor) RefundBet(opts *bind.TransactOpts, betId *big.Int) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "refundBet", betId)
}

// RefundBet is a paid mutator transaction binding the contract method 0xe1fdb4b4.
//
// Solidity: function refundBet(uint256 betId) returns()
func (_Betting *BettingSession) RefundBet(betId *big.Int) (*types.Transaction, error) {
	return _Betting.Contract.RefundBet(&_Betting.TransactOpts, betId)
}

// RefundBet is a paid mutator transaction binding the contract method 0xe1fdb4b4.
//
// Solidity: function refundBet(uint256 betId) returns()
func (_Betting *BettingTransactorSession) RefundBet(betId *big.Int) (*types.Transaction, error) {
	return _Betting.Contract.RefundBet(&_Betting.TransactOpts, betId)
}

// ResolveBet is a paid mutator transaction binding the contract method 0x4c36c36e.
//
// Solidity: function resolveBet(uint256 _betID) returns()
func (_Betting *BettingTransactor) ResolveBet(opts *bind.TransactOpts, _betID *big.Int) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "resolveBet", _betID)
}

// ResolveBet is a paid mutator transaction binding the contract method 0x4c36c36e.
//
// Solidity: function resolveBet(uint256 _betID) returns()
func (_Betting *BettingSession) ResolveBet(_betID *big.Int) (*types.Transaction, error) {
	return _Betting.Contract.ResolveBet(&_Betting.TransactOpts, _betID)
}

// ResolveBet is a paid mutator transaction binding the contract method 0x4c36c36e.
//
// Solidity: function resolveBet(uint256 _betID) returns()
func (_Betting *BettingTransactorSession) ResolveBet(_betID *big.Int) (*types.Transaction, error) {
	return _Betting.Contract.ResolveBet(&_Betting.TransactOpts, _betID)
}

// SetBotAddress is a paid mutator transaction binding the contract method 0x2d4f40c6.
//
// Solidity: function setBotAddress(address _botAddress) returns()
func (_Betting *BettingTransactor) SetBotAddress(opts *bind.TransactOpts, _botAddress common.Address) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "setBotAddress", _botAddress)
}

// SetBotAddress is a paid mutator transaction binding the contract method 0x2d4f40c6.
//
// Solidity: function setBotAddress(address _botAddress) returns()
func (_Betting *BettingSession) SetBotAddress(_botAddress common.Address) (*types.Transaction, error) {
	return _Betting.Contract.SetBotAddress(&_Betting.TransactOpts, _botAddress)
}

// SetBotAddress is a paid mutator transaction binding the contract method 0x2d4f40c6.
//
// Solidity: function setBotAddress(address _botAddress) returns()
func (_Betting *BettingTransactorSession) SetBotAddress(_botAddress common.Address) (*types.Transaction, error) {
	return _Betting.Contract.SetBotAddress(&_Betting.TransactOpts, _botAddress)
}

// SetBotFeeBasisPoints is a paid mutator transaction binding the contract method 0xb95964d5.
//
// Solidity: function setBotFeeBasisPoints(uint256 _botFeeBasisPoints) returns()
func (_Betting *BettingTransactor) SetBotFeeBasisPoints(opts *bind.TransactOpts, _botFeeBasisPoints *big.Int) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "setBotFeeBasisPoints", _botFeeBasisPoints)
}

// SetBotFeeBasisPoints is a paid mutator transaction binding the contract method 0xb95964d5.
//
// Solidity: function setBotFeeBasisPoints(uint256 _botFeeBasisPoints) returns()
func (_Betting *BettingSession) SetBotFeeBasisPoints(_botFeeBasisPoints *big.Int) (*types.Transaction, error) {
	return _Betting.Contract.SetBotFeeBasisPoints(&_Betting.TransactOpts, _botFeeBasisPoints)
}

// SetBotFeeBasisPoints is a paid mutator transaction binding the contract method 0xb95964d5.
//
// Solidity: function setBotFeeBasisPoints(uint256 _botFeeBasisPoints) returns()
func (_Betting *BettingTransactorSession) SetBotFeeBasisPoints(_botFeeBasisPoints *big.Int) (*types.Transaction, error) {
	return _Betting.Contract.SetBotFeeBasisPoints(&_Betting.TransactOpts, _botFeeBasisPoints)
}

// SetPriceFeedAddress is a paid mutator transaction binding the contract method 0x00cf5db4.
//
// Solidity: function setPriceFeedAddress(address _feedAddress) returns()
func (_Betting *BettingTransactor) SetPriceFeedAddress(opts *bind.TransactOpts, _feedAddress common.Address) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "setPriceFeedAddress", _feedAddress)
}

// SetPriceFeedAddress is a paid mutator transaction binding the contract method 0x00cf5db4.
//
// Solidity: function setPriceFeedAddress(address _feedAddress) returns()
func (_Betting *BettingSession) SetPriceFeedAddress(_feedAddress common.Address) (*types.Transaction, error) {
	return _Betting.Contract.SetPriceFeedAddress(&_Betting.TransactOpts, _feedAddress)
}

// SetPriceFeedAddress is a paid mutator transaction binding the contract method 0x00cf5db4.
//
// Solidity: function setPriceFeedAddress(address _feedAddress) returns()
func (_Betting *BettingTransactorSession) SetPriceFeedAddress(_feedAddress common.Address) (*types.Transaction, error) {
	return _Betting.Contract.SetPriceFeedAddress(&_Betting.TransactOpts, _feedAddress)
}

// SetUSDCAddress is a paid mutator transaction binding the contract method 0xaaf5bfc3.
//
// Solidity: function setUSDCAddress(address _usdcAddress) returns()
func (_Betting *BettingTransactor) SetUSDCAddress(opts *bind.TransactOpts, _usdcAddress common.Address) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "setUSDCAddress", _usdcAddress)
}

// SetUSDCAddress is a paid mutator transaction binding the contract method 0xaaf5bfc3.
//
// Solidity: function setUSDCAddress(address _usdcAddress) returns()
func (_Betting *BettingSession) SetUSDCAddress(_usdcAddress common.Address) (*types.Transaction, error) {
	return _Betting.Contract.SetUSDCAddress(&_Betting.TransactOpts, _usdcAddress)
}

// SetUSDCAddress is a paid mutator transaction binding the contract method 0xaaf5bfc3.
//
// Solidity: function setUSDCAddress(address _usdcAddress) returns()
func (_Betting *BettingTransactorSession) SetUSDCAddress(_usdcAddress common.Address) (*types.Transaction, error) {
	return _Betting.Contract.SetUSDCAddress(&_Betting.TransactOpts, _usdcAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Betting *BettingTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Betting *BettingSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Betting.Contract.TransferOwnership(&_Betting.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Betting *BettingTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Betting.Contract.TransferOwnership(&_Betting.TransactOpts, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _betID) returns()
func (_Betting *BettingTransactor) Withdraw(opts *bind.TransactOpts, _betID *big.Int) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "withdraw", _betID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _betID) returns()
func (_Betting *BettingSession) Withdraw(_betID *big.Int) (*types.Transaction, error) {
	return _Betting.Contract.Withdraw(&_Betting.TransactOpts, _betID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _betID) returns()
func (_Betting *BettingTransactorSession) Withdraw(_betID *big.Int) (*types.Transaction, error) {
	return _Betting.Contract.Withdraw(&_Betting.TransactOpts, _betID)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Betting *BettingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Betting *BettingSession) Receive() (*types.Transaction, error) {
	return _Betting.Contract.Receive(&_Betting.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Betting *BettingTransactorSession) Receive() (*types.Transaction, error) {
	return _Betting.Contract.Receive(&_Betting.TransactOpts)
}

// BettingBetActivedIterator is returned from FilterBetActived and is used to iterate over the raw logs and unpacked data for BetActived events raised by the Betting contract.
type BettingBetActivedIterator struct {
	Event *BettingBetActived // Event containing the contract specifics and raw log

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
func (it *BettingBetActivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BettingBetActived)
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
		it.Event = new(BettingBetActived)
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
func (it *BettingBetActivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BettingBetActivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BettingBetActived represents a BetActived event raised by the Betting contract.
type BettingBetActived struct {
	BetID        *big.Int
	UserB        common.Address
	OpeningPrice *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBetActived is a free log retrieval operation binding the contract event 0xd2540f7cf9cb3b638f442bcc1539e48c391249f4363c13006702f95356798ff8.
//
// Solidity: event BetActived(uint256 betID, address indexed userB, uint256 openingPrice)
func (_Betting *BettingFilterer) FilterBetActived(opts *bind.FilterOpts, userB []common.Address) (*BettingBetActivedIterator, error) {

	var userBRule []interface{}
	for _, userBItem := range userB {
		userBRule = append(userBRule, userBItem)
	}

	logs, sub, err := _Betting.contract.FilterLogs(opts, "BetActived", userBRule)
	if err != nil {
		return nil, err
	}
	return &BettingBetActivedIterator{contract: _Betting.contract, event: "BetActived", logs: logs, sub: sub}, nil
}

// WatchBetActived is a free log subscription operation binding the contract event 0xd2540f7cf9cb3b638f442bcc1539e48c391249f4363c13006702f95356798ff8.
//
// Solidity: event BetActived(uint256 betID, address indexed userB, uint256 openingPrice)
func (_Betting *BettingFilterer) WatchBetActived(opts *bind.WatchOpts, sink chan<- *BettingBetActived, userB []common.Address) (event.Subscription, error) {

	var userBRule []interface{}
	for _, userBItem := range userB {
		userBRule = append(userBRule, userBItem)
	}

	logs, sub, err := _Betting.contract.WatchLogs(opts, "BetActived", userBRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BettingBetActived)
				if err := _Betting.contract.UnpackLog(event, "BetActived", log); err != nil {
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

// ParseBetActived is a log parse operation binding the contract event 0xd2540f7cf9cb3b638f442bcc1539e48c391249f4363c13006702f95356798ff8.
//
// Solidity: event BetActived(uint256 betID, address indexed userB, uint256 openingPrice)
func (_Betting *BettingFilterer) ParseBetActived(log types.Log) (*BettingBetActived, error) {
	event := new(BettingBetActived)
	if err := _Betting.contract.UnpackLog(event, "BetActived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BettingBetClosedIterator is returned from FilterBetClosed and is used to iterate over the raw logs and unpacked data for BetClosed events raised by the Betting contract.
type BettingBetClosedIterator struct {
	Event *BettingBetClosed // Event containing the contract specifics and raw log

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
func (it *BettingBetClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BettingBetClosed)
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
		it.Event = new(BettingBetClosed)
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
func (it *BettingBetClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BettingBetClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BettingBetClosed represents a BetClosed event raised by the Betting contract.
type BettingBetClosed struct {
	BetID        *big.Int
	Winner       common.Address
	ClosingPrice *big.Int
	WinnerReward *big.Int
	FeeAmount    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBetClosed is a free log retrieval operation binding the contract event 0x1467f9463becd913c51269e1bd12fa890fc94406a5b940f86abb8a8bf34d9887.
//
// Solidity: event BetClosed(uint256 betID, address indexed winner, uint256 closingPrice, uint256 winnerReward, uint256 feeAmount)
func (_Betting *BettingFilterer) FilterBetClosed(opts *bind.FilterOpts, winner []common.Address) (*BettingBetClosedIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _Betting.contract.FilterLogs(opts, "BetClosed", winnerRule)
	if err != nil {
		return nil, err
	}
	return &BettingBetClosedIterator{contract: _Betting.contract, event: "BetClosed", logs: logs, sub: sub}, nil
}

// WatchBetClosed is a free log subscription operation binding the contract event 0x1467f9463becd913c51269e1bd12fa890fc94406a5b940f86abb8a8bf34d9887.
//
// Solidity: event BetClosed(uint256 betID, address indexed winner, uint256 closingPrice, uint256 winnerReward, uint256 feeAmount)
func (_Betting *BettingFilterer) WatchBetClosed(opts *bind.WatchOpts, sink chan<- *BettingBetClosed, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _Betting.contract.WatchLogs(opts, "BetClosed", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BettingBetClosed)
				if err := _Betting.contract.UnpackLog(event, "BetClosed", log); err != nil {
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

// ParseBetClosed is a log parse operation binding the contract event 0x1467f9463becd913c51269e1bd12fa890fc94406a5b940f86abb8a8bf34d9887.
//
// Solidity: event BetClosed(uint256 betID, address indexed winner, uint256 closingPrice, uint256 winnerReward, uint256 feeAmount)
func (_Betting *BettingFilterer) ParseBetClosed(log types.Log) (*BettingBetClosed, error) {
	event := new(BettingBetClosed)
	if err := _Betting.contract.UnpackLog(event, "BetClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BettingBetRefundedIterator is returned from FilterBetRefunded and is used to iterate over the raw logs and unpacked data for BetRefunded events raised by the Betting contract.
type BettingBetRefundedIterator struct {
	Event *BettingBetRefunded // Event containing the contract specifics and raw log

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
func (it *BettingBetRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BettingBetRefunded)
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
		it.Event = new(BettingBetRefunded)
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
func (it *BettingBetRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BettingBetRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BettingBetRefunded represents a BetRefunded event raised by the Betting contract.
type BettingBetRefunded struct {
	BetId  *big.Int
	UserA  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBetRefunded is a free log retrieval operation binding the contract event 0xf22517a5eae9f54be712c3150eca5ba925d233687b223c2ebf89272f12ee0865.
//
// Solidity: event BetRefunded(uint256 indexed betId, address userA, uint256 amount)
func (_Betting *BettingFilterer) FilterBetRefunded(opts *bind.FilterOpts, betId []*big.Int) (*BettingBetRefundedIterator, error) {

	var betIdRule []interface{}
	for _, betIdItem := range betId {
		betIdRule = append(betIdRule, betIdItem)
	}

	logs, sub, err := _Betting.contract.FilterLogs(opts, "BetRefunded", betIdRule)
	if err != nil {
		return nil, err
	}
	return &BettingBetRefundedIterator{contract: _Betting.contract, event: "BetRefunded", logs: logs, sub: sub}, nil
}

// WatchBetRefunded is a free log subscription operation binding the contract event 0xf22517a5eae9f54be712c3150eca5ba925d233687b223c2ebf89272f12ee0865.
//
// Solidity: event BetRefunded(uint256 indexed betId, address userA, uint256 amount)
func (_Betting *BettingFilterer) WatchBetRefunded(opts *bind.WatchOpts, sink chan<- *BettingBetRefunded, betId []*big.Int) (event.Subscription, error) {

	var betIdRule []interface{}
	for _, betIdItem := range betId {
		betIdRule = append(betIdRule, betIdItem)
	}

	logs, sub, err := _Betting.contract.WatchLogs(opts, "BetRefunded", betIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BettingBetRefunded)
				if err := _Betting.contract.UnpackLog(event, "BetRefunded", log); err != nil {
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

// ParseBetRefunded is a log parse operation binding the contract event 0xf22517a5eae9f54be712c3150eca5ba925d233687b223c2ebf89272f12ee0865.
//
// Solidity: event BetRefunded(uint256 indexed betId, address userA, uint256 amount)
func (_Betting *BettingFilterer) ParseBetRefunded(log types.Log) (*BettingBetRefunded, error) {
	event := new(BettingBetRefunded)
	if err := _Betting.contract.UnpackLog(event, "BetRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BettingNewBetCreatedIterator is returned from FilterNewBetCreated and is used to iterate over the raw logs and unpacked data for NewBetCreated events raised by the Betting contract.
type BettingNewBetCreatedIterator struct {
	Event *BettingNewBetCreated // Event containing the contract specifics and raw log

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
func (it *BettingNewBetCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BettingNewBetCreated)
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
		it.Event = new(BettingNewBetCreated)
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
func (it *BettingNewBetCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BettingNewBetCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BettingNewBetCreated represents a NewBetCreated event raised by the Betting contract.
type BettingNewBetCreated struct {
	BetID       *big.Int
	UserA       common.Address
	Amount      *big.Int
	IsLong      bool
	CreateTime  *big.Int
	ExpireTime  *big.Int
	ClosingTime *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewBetCreated is a free log retrieval operation binding the contract event 0xf74d83efdd7275db73dde45ff15fed93e1eb2ecd96c7f9bbb109fb73876ebb0c.
//
// Solidity: event NewBetCreated(uint256 betID, address indexed userA, uint256 amount, bool isLong, uint256 createTime, uint256 expireTime, uint256 closingTime)
func (_Betting *BettingFilterer) FilterNewBetCreated(opts *bind.FilterOpts, userA []common.Address) (*BettingNewBetCreatedIterator, error) {

	var userARule []interface{}
	for _, userAItem := range userA {
		userARule = append(userARule, userAItem)
	}

	logs, sub, err := _Betting.contract.FilterLogs(opts, "NewBetCreated", userARule)
	if err != nil {
		return nil, err
	}
	return &BettingNewBetCreatedIterator{contract: _Betting.contract, event: "NewBetCreated", logs: logs, sub: sub}, nil
}

// WatchNewBetCreated is a free log subscription operation binding the contract event 0xf74d83efdd7275db73dde45ff15fed93e1eb2ecd96c7f9bbb109fb73876ebb0c.
//
// Solidity: event NewBetCreated(uint256 betID, address indexed userA, uint256 amount, bool isLong, uint256 createTime, uint256 expireTime, uint256 closingTime)
func (_Betting *BettingFilterer) WatchNewBetCreated(opts *bind.WatchOpts, sink chan<- *BettingNewBetCreated, userA []common.Address) (event.Subscription, error) {

	var userARule []interface{}
	for _, userAItem := range userA {
		userARule = append(userARule, userAItem)
	}

	logs, sub, err := _Betting.contract.WatchLogs(opts, "NewBetCreated", userARule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BettingNewBetCreated)
				if err := _Betting.contract.UnpackLog(event, "NewBetCreated", log); err != nil {
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

// ParseNewBetCreated is a log parse operation binding the contract event 0xf74d83efdd7275db73dde45ff15fed93e1eb2ecd96c7f9bbb109fb73876ebb0c.
//
// Solidity: event NewBetCreated(uint256 betID, address indexed userA, uint256 amount, bool isLong, uint256 createTime, uint256 expireTime, uint256 closingTime)
func (_Betting *BettingFilterer) ParseNewBetCreated(log types.Log) (*BettingNewBetCreated, error) {
	event := new(BettingNewBetCreated)
	if err := _Betting.contract.UnpackLog(event, "NewBetCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BettingOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the Betting contract.
type BettingOwnershipTransferRequestedIterator struct {
	Event *BettingOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *BettingOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BettingOwnershipTransferRequested)
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
		it.Event = new(BettingOwnershipTransferRequested)
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
func (it *BettingOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BettingOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BettingOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the Betting contract.
type BettingOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Betting *BettingFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BettingOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Betting.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BettingOwnershipTransferRequestedIterator{contract: _Betting.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Betting *BettingFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BettingOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Betting.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BettingOwnershipTransferRequested)
				if err := _Betting.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Betting *BettingFilterer) ParseOwnershipTransferRequested(log types.Log) (*BettingOwnershipTransferRequested, error) {
	event := new(BettingOwnershipTransferRequested)
	if err := _Betting.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BettingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Betting contract.
type BettingOwnershipTransferredIterator struct {
	Event *BettingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BettingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BettingOwnershipTransferred)
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
		it.Event = new(BettingOwnershipTransferred)
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
func (it *BettingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BettingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BettingOwnershipTransferred represents a OwnershipTransferred event raised by the Betting contract.
type BettingOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Betting *BettingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BettingOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Betting.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BettingOwnershipTransferredIterator{contract: _Betting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Betting *BettingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BettingOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Betting.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BettingOwnershipTransferred)
				if err := _Betting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Betting *BettingFilterer) ParseOwnershipTransferred(log types.Log) (*BettingOwnershipTransferred, error) {
	event := new(BettingOwnershipTransferred)
	if err := _Betting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BettingRewardWithdrawalIterator is returned from FilterRewardWithdrawal and is used to iterate over the raw logs and unpacked data for RewardWithdrawal events raised by the Betting contract.
type BettingRewardWithdrawalIterator struct {
	Event *BettingRewardWithdrawal // Event containing the contract specifics and raw log

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
func (it *BettingRewardWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BettingRewardWithdrawal)
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
		it.Event = new(BettingRewardWithdrawal)
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
func (it *BettingRewardWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BettingRewardWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BettingRewardWithdrawal represents a RewardWithdrawal event raised by the Betting contract.
type BettingRewardWithdrawal struct {
	BetID  *big.Int
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardWithdrawal is a free log retrieval operation binding the contract event 0xc5125b27e7e5dd75644ec565391448b5229c80f3b00d5f3b032fe7b263e7c031.
//
// Solidity: event RewardWithdrawal(uint256 betID, address indexed winner, uint256 amount)
func (_Betting *BettingFilterer) FilterRewardWithdrawal(opts *bind.FilterOpts, winner []common.Address) (*BettingRewardWithdrawalIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _Betting.contract.FilterLogs(opts, "RewardWithdrawal", winnerRule)
	if err != nil {
		return nil, err
	}
	return &BettingRewardWithdrawalIterator{contract: _Betting.contract, event: "RewardWithdrawal", logs: logs, sub: sub}, nil
}

// WatchRewardWithdrawal is a free log subscription operation binding the contract event 0xc5125b27e7e5dd75644ec565391448b5229c80f3b00d5f3b032fe7b263e7c031.
//
// Solidity: event RewardWithdrawal(uint256 betID, address indexed winner, uint256 amount)
func (_Betting *BettingFilterer) WatchRewardWithdrawal(opts *bind.WatchOpts, sink chan<- *BettingRewardWithdrawal, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _Betting.contract.WatchLogs(opts, "RewardWithdrawal", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BettingRewardWithdrawal)
				if err := _Betting.contract.UnpackLog(event, "RewardWithdrawal", log); err != nil {
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

// ParseRewardWithdrawal is a log parse operation binding the contract event 0xc5125b27e7e5dd75644ec565391448b5229c80f3b00d5f3b032fe7b263e7c031.
//
// Solidity: event RewardWithdrawal(uint256 betID, address indexed winner, uint256 amount)
func (_Betting *BettingFilterer) ParseRewardWithdrawal(log types.Log) (*BettingRewardWithdrawal, error) {
	event := new(BettingRewardWithdrawal)
	if err := _Betting.contract.UnpackLog(event, "RewardWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
