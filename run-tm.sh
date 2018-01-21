echo tendermint start
tendermint init
tendermint unsafe_reset_all
tendermint node --consensus.create_empty_blocks=false
echo tendermint end