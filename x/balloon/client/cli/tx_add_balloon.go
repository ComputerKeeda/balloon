package cli

import (
	"strconv"

	"github.com/ComputerKeeda/balloon/x/balloon/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddBalloon() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-balloon [balloon-name] [balloon-height]",
		Short: "Broadcast message addBalloon",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBalloonName := args[0]
			argBalloonHeight := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddBalloon(
				clientCtx.GetFromAddress().String(),
				argBalloonName,
				argBalloonHeight,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
