package types

import errorsmod "cosmossdk.io/errors"

const MAX_ALIAS_LEN = 25

func checkAliasLimit(alias string) error {
	if len(alias) > MAX_ALIAS_LEN {
		return errorsmod.Wrapf(ErrInvalidData, "alias is too long (%d) maximum is %d", len(alias), MAX_ALIAS_LEN)
	}

	return nil
}
