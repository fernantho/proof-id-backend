VALUE=1
if [ "$1" ]; then
    VALUE=$1
fi

echo "[" ${VALUE} "]" > public.json

echo "Verify the proof..."
# Verify the proof
snarkjs groth16 verify verification_key.json public.json proof.json
