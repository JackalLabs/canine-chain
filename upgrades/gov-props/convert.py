def main():
    f = open("v2.0.0.md", "r")
    mdown = f.read()

    print(mdown.replace("\n", '\\n'))


main()