# MAX_SUPPLY

```solidity
uint256 public constant MAX_SUPPLY = 1_000_000 * 10**18;
```

## Breakdown

uint256 → Declares an unsigned integer (a number ≥ 0)

public constant → The value is publicly accessible and unchangeable after deployment.

MAX_SUPPLY → Name of the constant, indicating the total maximum number of tokens that can ever exist.

1_000_000 → Represents 1 million tokens.

10**18 → Since ERC-20 tokens typically use 18 decimal places, multiplying by 10^18 ensures precision.

Example: If 1 token = 1 unit, it wouldn’t support fractions (like 0.5 tokens).

By using 10^18, 1 token is represented as 1 *10¹⁸ smallest units, allowing fractional values (e.g., 0.5 tokens = 5* 10^17).

Final Value Stored:

MAX_SUPPLY = 1,000,000 * 1,000,000,000,000,000,000
= 1,000,000,000,000,000,000,000,000 (1 million tokens with 18 decimals).

Purpose:
✅ Prevents minting beyond MAX_SUPPLY.
✅ Ensures controlled token issuance.
✅ Matches ERC-20 precision standards (18 decimal places).

Let me know if you want to modify it (e.g., fewer decimals or different supply limits).
