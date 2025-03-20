# Variables
ABIGEN=abigen
CONTRACTS_DIR=src
OUTPUT_DIR=bindings

# Phony targets (not actual files)
.PHONY: all clean

# Default target
all: generate-ecokash generate-memekash

# Rule: Compile EcoKash.sol to generate artifacts for all contracts
$(CONTRACTS_DIR)/EcoKash.abi: $(CONTRACTS_DIR)/EcoKash.sol
	solc --abi --bin -o $(CONTRACTS_DIR) $(CONTRACTS_DIR)/EcoKash.sol

# Rule: Generate Go bindings for EcoKash
generate-ecokash: $(CONTRACTS_DIR)/EcoKash.abi
	$(ABIGEN) --abi $(CONTRACTS_DIR)/EcoKash.abi --bin $(CONTRACTS_DIR)/EcoKash.bin --pkg ecokash --out $(OUTPUT_DIR)/ecokash.go

# Rule: Generate Go bindings for MemeKash
generate-memekash: $(CONTRACTS_DIR)/EcoKash.abi
	$(ABIGEN) --abi $(CONTRACTS_DIR)/MemeKash.abi --bin $(CONTRACTS_DIR)/MemeKash.bin --pkg memekash --out $(OUTPUT_DIR)/memekash.go

# Clean up generated files
clean:
	rm -rf $(OUTPUT_DIR)/*.go
	rm -rf $(CONTRACTS_DIR)/*.abi $(CONTRACTS_DIR)/*.bin




# Target: The dish you want to make (e.g., EcoKash.abi).
# Prerequisite: The ingredients you need (e.g., EcoKash.sol).
# Command: The steps to cook the dish (e.g., run solc).