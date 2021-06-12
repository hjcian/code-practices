def group_by_owners(files: dict):
    ret = {}
    for filename, name in files.items():
        fileList = ret.get(name, [])
        fileList.append(filename)
        ret[name] = fileList
    return ret


if __name__ == "__main__":
    files = {
        'Input.txt': 'Randy',
        'Code.py': 'Stan',
        'Output.txt': 'Randy'
    }

    print(group_by_owners(files))
    ans = {'Randy': ['Input.txt', 'Output.txt'], 'Stan': ['Code.py']}
