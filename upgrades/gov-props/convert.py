import sys

def main():
    args = sys.argv
    if len(args) == 1 :
        print("needs file name!")
        return
    f = open(args[1], "r")
    mdown = f.read()

    print(mdown.replace("\n", '\\n'))

if __name__ == "__main__":
    main()