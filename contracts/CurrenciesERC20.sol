//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


//import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "./interfaces/ICurrenciesERC20.sol";

//../node_modules/
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";

/**
 *      CurrenciesERC20
 * @title CurrenciesERC20
 * @author JackBekket https://github.com/JackBekket
 * @dev This contract allow to use erc20 tokens as a currency in crowdsale-like contracts
 *
 */
contract CurrenciesERC20 is ReentrancyGuard, Ownable, ERC165 {
    using SafeMath for uint256;
    //  using SafeERC20 for IERC20;

    // Interface to currency token
    //IERC20 public _currency_token;

    // Supported erc20 currencies: .. to be extended.  This is hard-coded values
    /**
     * @dev Hardcoded (not-extetiable after deploy) erc20 currencies
     */
    enum CurrencyERC20 {
        USDT,
        USDC,
        DAI,
        WETH,
        WBTC,
        VXPPL
    }

    struct CurrencyERC20_Custom {
        address contract_address;
        IERC20Metadata itoken; // contract interface
    }

    // map currency contract addresses
    mapping(CurrencyERC20 => IERC20Metadata) public _currencies_hardcoded; // should be internal?

    // mapping from name to currency contract (protected)
    mapping(string => CurrencyERC20_Custom) public _currencies_custom;

    // mapping from name to currency contract defined by users (not protected against scum)
    mapping(string => CurrencyERC20_Custom) public _currencies_custom_user;

    
    bytes4 private _INTERFACE_ID_CURRECIES = 0x033a36bd;

    function AddCustomCurrency(address _token_contract) public {
        IERC20Metadata _currency_contract = IERC20Metadata(_token_contract);

        // if (_currency_contract.name != '0x0')

        string memory _name_c = _currency_contract.name(); // @note -- some contracts just have name as public string, but do not have name() function!!! see difference between 0.4.0 and 0.8.0 OZ standarts need future consideration
        //  uint8 _dec = _currency_contract.decimals();

        address _owner_c = owner();
        if (msg.sender == _owner_c) {
            require(
                _currencies_custom[_name_c].contract_address == address(0),
                "AddCustomCurrency[admin]: Currency token contract with this address is already exists"
            );
            _currencies_custom[_name_c].itoken = _currency_contract;
            //   _currencies_custom[_name_c].decimals = _dec;
            _currencies_custom[_name_c].contract_address = _token_contract;
        } else {
            require(
                _currencies_custom_user[_name_c].contract_address == address(0),
                "AddCustomCurrency[user]: Currency token contract with this address is already exists"
            );
            _currencies_custom_user[_name_c].itoken = _currency_contract;
            //  _currencies_custom_user[_name_c].decimals = _dec;
            _currencies_custom_user[_name_c].contract_address = _token_contract;
        }
    }

    constructor(
        address US_Tether,
        address US_Circle,
        address DAI,
        address W_Ethereum,
        address WBTC,
        address VXPPL
    ) {
        require(US_Tether != address(0), "USDT contract address is zero!");
        require(US_Circle != address(0), "US_Circle contract address is zero!");
        require(DAI != address(0), "DAI contract address is zero!");
        require(
            W_Ethereum != address(0),
            "W_Ethereum contract address is zero!"
        );
        require(WBTC != address(0), "WBTC contract address is zero!");

        _currencies_hardcoded[CurrencyERC20.USDT] = IERC20Metadata(US_Tether);
        _currencies_hardcoded[CurrencyERC20.USDT] == IERC20Metadata(US_Tether);
        _currencies_hardcoded[CurrencyERC20.USDC] = IERC20Metadata(US_Circle);
        _currencies_hardcoded[CurrencyERC20.DAI] = IERC20Metadata(DAI);
        _currencies_hardcoded[CurrencyERC20.WETH] = IERC20Metadata(W_Ethereum);
        _currencies_hardcoded[CurrencyERC20.WBTC] = IERC20Metadata(WBTC);
        _currencies_hardcoded[CurrencyERC20.VXPPL] = IERC20Metadata(VXPPL);


    }

    function get_hardcoded_currency(CurrencyERC20 currency)
        public
        view
        returns (IERC20Metadata)
    {
        return _currencies_hardcoded[currency];
    }

    function supportsInterface(bytes4 interfaceId)
    public view override
    returns (bool) {
       return interfaceId == type(ICurrenciesERC20).interfaceId || super.supportsInterface(interfaceId);
    }
}
