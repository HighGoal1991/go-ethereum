package ethchain

import (
	"fmt"
	"github.com/ethereum/eth-go/ethutil"
	"math/big"
	"strings"
)

type Code []byte

func (self Code) String() string {
	return strings.Join(Disassemble(self), " ")
}

type StateObject struct {
	// Address of the object
	address []byte
	// Shared attributes
	Amount     *big.Int
	ScriptHash []byte
	Nonce      uint64
	// Contract related attributes
	state      *State
	script     Code
	initScript Code

	// Total gas pool is the total amount of gas currently
	// left if this object is the coinbase. Gas is directly
	// purchased of the coinbase.
	gasPool *big.Int
}

// Converts an transaction in to a state object
func MakeContract(tx *Transaction, state *State) *StateObject {
	// Create contract if there's no recipient
	if tx.IsContract() {
		addr := tx.CreationAddress()

		contract := state.NewStateObject(addr)
		contract.initScript = tx.Data
		contract.state = NewState(ethutil.NewTrie(ethutil.Config.Db, ""))

		return contract
	}

	return nil
}

func NewStateObject(addr []byte) *StateObject {
	return &StateObject{address: addr, Amount: new(big.Int), gasPool: new(big.Int)}
}

func NewContract(address []byte, Amount *big.Int, root []byte) *StateObject {
	contract := &StateObject{address: address, Amount: Amount, Nonce: 0}
	contract.state = NewState(ethutil.NewTrie(ethutil.Config.Db, string(root)))

	return contract
}

// Returns a newly created account
func NewAccount(address []byte, amount *big.Int) *StateObject {
	account := &StateObject{address: address, Amount: amount, Nonce: 0}

	return account
}

func NewStateObjectFromBytes(address, data []byte) *StateObject {
	object := &StateObject{address: address}
	object.RlpDecode(data)

	return object
}

func (c *StateObject) State() *State {
	return c.state
}

func (c *StateObject) N() *big.Int {
	return big.NewInt(int64(c.Nonce))
}

func (c *StateObject) Addr(addr []byte) *ethutil.Value {
	return ethutil.NewValueFromBytes([]byte(c.state.trie.Get(string(addr))))
}

func (c *StateObject) SetAddr(addr []byte, value interface{}) {
	c.state.trie.Update(string(addr), string(ethutil.NewValue(value).Encode()))
}

func (c *StateObject) SetStorage(num *big.Int, val *ethutil.Value) {
	addr := ethutil.BigToBytes(num, 256)

	// FIXME This should be handled in the Trie it self
	if val.BigInt().Cmp(ethutil.Big0) == 0 {
		c.state.trie.Delete(string(addr))

		return
	}

	c.SetAddr(addr, val)
}

func (c *StateObject) GetStorage(num *big.Int) *ethutil.Value {
	nb := ethutil.BigToBytes(num, 256)

	return c.Addr(nb)
}

/* DEPRECATED */
func (c *StateObject) GetMem(num *big.Int) *ethutil.Value {
	return c.GetStorage(num)
}

func (c *StateObject) GetInstr(pc *big.Int) *ethutil.Value {
	if int64(len(c.script)-1) < pc.Int64() {
		return ethutil.NewValue(0)
	}

	return ethutil.NewValueFromBytes([]byte{c.script[pc.Int64()]})
}

// Return the gas back to the origin. Used by the Virtual machine or Closures
func (c *StateObject) ReturnGas(gas, price *big.Int, state *State) {
	/*
		remainder := new(big.Int).Mul(gas, price)
		c.AddAmount(remainder)
	*/
}

func (c *StateObject) AddAmount(amount *big.Int) {
	c.SetAmount(new(big.Int).Add(c.Amount, amount))

	ethutil.Config.Log.Printf(ethutil.LogLevelInfo, "%x: #%d %v (+ %v)\n", c.Address(), c.Nonce, c.Amount, amount)
}

func (c *StateObject) SubAmount(amount *big.Int) {
	c.SetAmount(new(big.Int).Sub(c.Amount, amount))

	ethutil.Config.Log.Printf(ethutil.LogLevelInfo, "%x: #%d %v (- %v)\n", c.Address(), c.Nonce, c.Amount, amount)
}

func (c *StateObject) SetAmount(amount *big.Int) {
	c.Amount = amount
}

func (c *StateObject) ConvertGas(gas, price *big.Int) error {
	total := new(big.Int).Mul(gas, price)
	if total.Cmp(c.Amount) > 0 {
		return fmt.Errorf("insufficient amount: %v, %v", c.Amount, total)
	}

	c.SubAmount(total)

	return nil
}

func (self *StateObject) SetGasPool(gasLimit *big.Int) {
	self.gasPool = new(big.Int).Set(gasLimit)

	ethutil.Config.Log.Printf(ethutil.LogLevelSystem, "%x: fuel (+ %v)", self.Address(), self.gasPool)
}

func (self *StateObject) BuyGas(gas, price *big.Int) error {
	if self.gasPool.Cmp(gas) < 0 {
		return GasLimitError(self.gasPool, gas)
	}

	rGas := new(big.Int).Set(gas)
	rGas.Mul(rGas, price)

	self.AddAmount(rGas)

	return nil
}

func (self *StateObject) RefundGas(gas, price *big.Int) {
	self.gasPool.Add(self.gasPool, gas)

	rGas := new(big.Int).Set(gas)
	rGas.Mul(rGas, price)

	self.Amount.Sub(self.Amount, rGas)
}

func (self *StateObject) Copy() *StateObject {
	stateObject := NewStateObject(self.Address())
	stateObject.Amount.Set(self.Amount)
	stateObject.ScriptHash = ethutil.CopyBytes(self.ScriptHash)
	stateObject.Nonce = self.Nonce
	if self.state != nil {
		stateObject.state = self.state.Copy()
	}
	stateObject.script = ethutil.CopyBytes(self.script)
	stateObject.initScript = ethutil.CopyBytes(self.initScript)
	//stateObject.gasPool.Set(self.gasPool)

	return self
}

func (self *StateObject) Set(stateObject *StateObject) {
	self = stateObject
}

/*
func (self *StateObject) Copy() *StateObject {
	stCopy := &StateObject{}
	stCopy.address = make([]byte, len(self.address))
	copy(stCopy.address, self.address)
	stCopy.Amount = new(big.Int).Set(self.Amount)
	stCopy.ScriptHash = make([]byte, len(self.ScriptHash))
	copy(stCopy.ScriptHash, self.ScriptHash)
	stCopy.Nonce = self.Nonce
	if self.state != nil {
		stCopy.state = self.state.Copy()
	}
	stCopy.script = make([]byte, len(self.script))
	copy(stCopy.script, self.script)
	stCopy.initScript = make([]byte, len(self.initScript))
	copy(stCopy.initScript, self.initScript)

	return stCopy
}
*/

// Returns the address of the contract/account
func (c *StateObject) Address() []byte {
	return c.address
}

// Returns the main script body
func (c *StateObject) Script() Code {
	return c.script
}

// Returns the initialization script
func (c *StateObject) Init() Code {
	return c.initScript
}

// State object encoding methods
func (c *StateObject) RlpEncode() []byte {
	var root interface{}
	if c.state != nil {
		root = c.state.trie.Root
	} else {
		root = ""
	}

	return ethutil.Encode([]interface{}{c.Nonce, c.Amount, root, ethutil.Sha3Bin(c.script)})
}

func (c *StateObject) RlpDecode(data []byte) {
	decoder := ethutil.NewValueFromBytes(data)

	c.Nonce = decoder.Get(0).Uint()
	c.Amount = decoder.Get(1).BigInt()
	c.state = NewState(ethutil.NewTrie(ethutil.Config.Db, decoder.Get(2).Interface()))

	c.ScriptHash = decoder.Get(3).Bytes()

	c.script, _ = ethutil.Config.Db.Get(c.ScriptHash)
}

// Storage change object. Used by the manifest for notifying changes to
// the sub channels.
type StorageState struct {
	StateAddress []byte
	Address      []byte
	Value        *big.Int
}
