def get_directions():
    inputFile = open("input.txt", "r")
    directions = inputFile.readlines()
    inputFile.close()
    return directions

def handle_move(facing, value, lat, lng):
    if facing == "N":
        return (lat + value, lng)
    elif facing == "S":
        return (lat - value, lng)
    elif facing == "E":
        return (lat, lng + value)
    else: # facing == "W"
        return (lat, lng - value)

def handle_turn(action, value, facing):
    compass = ["N", "E", "S", "W"]
    curIndex = 0

    if facing == "N":
        curIndex = 0
    elif facing == "S":
        curIndex = 2
    elif facing == "E":
        curIndex = 1
    else: # facing == "W"
        curIndex = 3
    
    numTurns = value / 90
    if action == "R":
        while curIndex + numTurns > 3:
            numTurns -= 4
        return compass[curIndex + numTurns]
    else: # action == "L"
        while curIndex - numTurns < 0:
            curIndex += 4
        return compass[curIndex - numTurns]

def handle_directions(directions):
    facing = "E"
    lat = 0 # N/S (+/-)
    lng = 0 # E/W (+/-)

    for direction in directions:
        action = direction[0]
        value = int(direction[1:])
        if action == "F":
            (lat, lng) = handle_move(facing, value, lat, lng)
        elif action == "L" or action == "R":
            facing = handle_turn(action, value, facing)
        else:
            (lat, lng) = handle_move(action, value, lat, lng)
        print (facing, lat, lng)
    return (lat, lng)

def main():
    directions = get_directions()
    (lat, lng) = handle_directions(directions)
    return abs(lat) + abs(lng) 

if __name__ == "__main__":
    print main()