from rembg import remove
import sys

input_path = sys.argv[1]
output_path = sys.argv[2]

with open(input_path, 'rb') as i:
    input_bytes = i.read()

output_bytes = remove(input_bytes)

with open(output_path, 'wb') as o:
    o.write(output_bytes)