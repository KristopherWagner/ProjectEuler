def get_adapters():
    inputFile = open("input.txt", "r")
    adapters = inputFile.readlines()
    inputFile.close()

    for i in xrange(0, len(adapters)):
        adapters[i] = int(adapters[i].strip())
    sortedAdapters = sorted([0] + adapters)
    sortedAdapters.reverse()
    return sortedAdapters

'''
my original solution which kind of works but also doesn't, check out the explanation at the end
def create_arrangement(adapters):
    if len(adapters) == 1:
        return 1
    else:
        arrangements = 0
        cur = adapters[0]
        i = 1
        while i < len(adapters):
            difference = adapters[i] - cur
            if difference > 0 and difference < 4:
                arrangements += create_arrangement(adapters[i:])
            i += 1
        return arrangements
'''

def get_optional_adapters(adapters):
    optionalAdapters = []
    for i in xrange(1, len(adapters) - 1):
        if adapters[i - 1] - adapters[i + 1] <= 3:
            optionalAdapters.append(adapters[i])
    return optionalAdapters

def calculate_arrangements(optionalAdapters):
    arrangements = 1
    for i in xrange(0, len(optionalAdapters)):
        try:
            optionalAdapters.index(optionalAdapters[i] + 1)
            optionalAdapters.index(optionalAdapters[i] + 2)
            arrangements += 3 * arrangements / 4
        except:
            arrangements *= 2 
    return arrangements

def main():
    adapters = get_adapters()
    optionalAdapters = get_optional_adapters(adapters)
    return calculate_arrangements(optionalAdapters)



if __name__ == "__main__":
    print main()

'''
Just a note, after 5 minutes of running, I wondered why my original solution didn't have an answer
even though it worked for both sample sets. Surprise, it's not efficient enough. For an
explanation, check out this link:
https://davidlozzi.com/2020/12/10/advent-of-code-day-10-check-back-in-629-days/
'''