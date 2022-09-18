pragma circom 2.0.0;

template RangeProof() {
    signal input salary;
    signal input times;
    signal input price;
    signal threshold;
    signal aux;
    signal output result;

    // Operations
    threshold <-- times*price;
    aux <-- threshold <= times*salary ? 10 : 20;

    // Constraints
    result <== aux*1 + 0; 
}

component main = RangeProof();