def main():
    f = open("v3.4.0.md", "r")
    mdown = f.read()

    print(mdown.replace("\n", '\\n'))


main()