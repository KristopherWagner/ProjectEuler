def get_directions():
    inputFile = open("input.txt", "r")
    directions = inputFile.readlines()
    inputFile.close()
    return directions

def handle_move_ship(ship, waypoint, value):
    for i in xrange(0, value):
        ship = (ship[0] + waypoint[0], ship[1] + waypoint[1])
    return ship 

def handle_move_waypoint(facing, waypoint, value):
    (lat, lng) = waypoint
    if facing == "N":
        return (lat + value, lng)
    elif facing == "S":
        return (lat - value, lng)
    elif facing == "E":
        return (lat, lng + value)
    else: # facing == "W"
        return (lat, lng - value)

def handle_rotate(waypoint, action, value):
    numTurns = value / 90
    while numTurns > 4:
        print numTurns
        numTurns -= 4

    if value != 90 and value != 180:
        print value

    changes = [(-1, 1), (1, -1)]
    iterator = 0
    if waypoint[0] > 0: # N
        if waypoint[1] > 0: # NE
            iterator = 0
        else: # NW
            iterator = 1
    else: # S
        if waypoint[1] > 0: # SE
            iterator = 1
        else: # SW
            iterator = 0
    
    if action == "L":
        iterator = (iterator + 1) % 2

    for i in xrange(0, numTurns):
        waypoint = (waypoint[0] * changes[iterator][0], waypoint[1] * changes[iterator][1])
        iterator = (iterator + 1) % 2
    return waypoint
    

def handle_directions(directions):
    facing = "E"
    ship = (0, 0) # (N/S, E/W) +/-
    waypoint = (1, 10) # (N/S, E/W) +/-

    for direction in directions:
        action = direction[0]
        value = int(direction[1:])
        if action == "F":
            ship = handle_move_ship(ship, waypoint, value)
        elif action == "L" or action == "R":
            waypoint = handle_rotate(waypoint, action, value)
        else:
            waypoint = handle_move_waypoint(action, waypoint, value)
    return ship

def main():
    directions = get_directions()
    (lat, lng) = handle_directions(directions)
    return abs(lat) + abs(lng) 

if __name__ == "__main__":
    print main()