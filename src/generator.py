import os
import yaml


Loader = getattr(yaml, 'CSafeLoader', yaml.SafeLoader)

ACTION_YML_FILE = os.env['INPUT_ACTION_FILE']
INPUT_README_FILE = os.env['INPUT_README']


def main() -> int:
    with open(ACTION_YML_FILE) as f:
        inputs = yaml.load(f, Loader=Loader)

    with open(INPUT_README_FILE) as f:
        contents = f.read()
    
    before, delim, _ = contents.partition('[generated-inputs]: # (generated)\n')
     
    ## TODO: Return a table of input and output
    rest = ""

    with open(INPUT_README_FILE, 'w') as f:
        f.write(before + delim + rest + '\n')

    return 0


if __name__ == '__main__':
    exit(main())
