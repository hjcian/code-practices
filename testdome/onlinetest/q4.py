def main():
    file = None
    first_line = None
    try:
        file = open("./aaa.txt")
        print("file", file)
        first_line = file.readline()
    except FileNotFoundError as e:
        print(e)
    finally:
        if file is not None:
            print(
                "[should not print this line] because the file is not opened successfully")
            try:
                file.close()
            except:
                pass
    return first_line


print(main())
