def get_seats():
    inputFile = open("input.txt", "r")
    seats = inputFile.readlines()
    inputFile.close()
    return seats

def print_seats(seats):
    for row in seats:
        print row.strip()

def get_num_occupied(seat, seats):
    (row, column) = seat
    numOccupied = 0
    for r in xrange(row - 1, row + 2):
        if r >= 0 and r < len(seats):
            for c in xrange(column - 1, column + 2):
                if c >= 0 and c < len(seats[r]):
                    if seats[r][c] == '#':
                        if (r, c) != (row, column):
                            numOccupied += 1
    return numOccupied

def execute_round(seats):
    newSeats = []
    r = 0
    while r < len(seats):
        c = 0
        newSeats.append('')
        while c < len(seats[r]):
            seat = (r, c)
            if seats[r][c] == 'L' and get_num_occupied(seat, seats) == 0:
                newSeats[r] += '#'
            elif seats[r][c] == '#' and get_num_occupied(seat, seats) >= 4:
                newSeats[r] += 'L'
            else:
                newSeats[r] += seats[r][c]
            c += 1
        r += 1
    return newSeats

def count_occupied_seats(seats):
    occupied = 0
    r = 0
    while r < len(seats):
        c = 0
        while c < len(seats[r]):
            if seats[r][c] == '#':
                occupied += 1
            c += 1
        r += 1
    return occupied

def main():
    seats = get_seats()
    lastSeats = []
    while lastSeats != seats:
        lastSeats = seats
        seats = execute_round(seats)
    return count_occupied_seats(seats)

if __name__ == "__main__":
    print main()