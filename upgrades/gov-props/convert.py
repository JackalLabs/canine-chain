def main():
    f = open("v4.0.0.md", "r")
    mdown = f.read()

    print(mdown.replace("\n", '\\n'))


main()
