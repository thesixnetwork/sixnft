rm -rf ~/.sixnft/data
rm -rf ~/.sixnft/wasm
rm ~/.sixnft/config/addrbook.json
mkdir ~/.sixnft/data/
touch ~/.sixnft/data/priv_validator_state.json
echo '{"height": "0", "round": 0,"step": 0}' > ~/.sixnft/data/priv_validator_state.json