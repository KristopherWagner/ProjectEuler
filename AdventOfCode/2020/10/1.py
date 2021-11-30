def get_adapters():
    inputFile = open("input.txt", "r")
    adapters = inputFile.readlines()
    inputFile.close()

    for i in xrange(0, len(adapters)):
        adapters[i] = int(adapters[i].strip())
    return sorted(adapters)

def main():
    adapters = get_adapters()
    jolts = 0
    ones = 0
    threes = 1
    for adapter in adapters:
        difference = adapter - jolts
        print (adapter, jolts, difference)
        if difference == 1:
            ones += 1
        elif difference == 3:
            threes += 1
        jolts = adapter
    print (1, ones)
    print (3, threes)
    return ones * threes

if __name__ == "__main__":
    print main()