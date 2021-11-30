def get_seats():
    inputFile = open("input.txt", "r")
    seats = inputFile.readlines()
    inputFile.close()
    return seats

def print_seats(seats):
    for row in seats:
        print row.strip()

def is_seat(seat):
    return seat == 'L' or seat == '#'

def check_north(seat, seats):
    (row, column) = seat
    r = row - 1
    found = False
    while r >= 0 and not found:
        found = is_seat(seats[r][column])
        r -= 1
    if found and seats[r + 1][column] == '#':
        return 1
    else:
        return 0

def check_northeast(seat, seats):
    (row, column) = seat
    r = row - 1
    c = column + 1
    found = False
    while r >= 0 and c < len(seats[r]) and not found:
        found = is_seat(seats[r][c])
        r -= 1
        c += 1
    if found and seats[r + 1][c - 1] == '#':
        return 1
    else:
        return 0

def check_east(seat, seats):
    (row, column) = seat
    c = column + 1
    found = False
    while c < len(seats[row]) and not found:
        found = is_seat(seats[row][c])
        c += 1
    if found and seats[row][c - 1] == '#':
        return 1
    else:
        return 0

def check_southeast(seat, seats):
    (row, column) = seat
    r = row + 1
    c = column + 1
    found = False
    while r < len(seats) and c < len(seats[r]) and not found:
        found = is_seat(seats[r][c])
        r += 1
        c += 1
    if found and seats[r - 1][c - 1] == '#':
        return 1
    else:
        return 0

def check_south(seat, seats):
    (row, column) = seat
    r = row + 1
    found = False
    while r < len(seats) and not found:
        found = is_seat(seats[r][column])
        r += 1
    if found and seats[r - 1][column] == '#':
        return 1
    else:
        return 0

def check_southwest(seat, seats):
    (row, column) = seat
    r = row + 1
    c = column - 1
    found = False
    while r < len(seats) and c >= 0 and not found:
        found = is_seat(seats[r][c])
        r += 1
        c -= 1
    if found and seats[r - 1][c + 1] == '#':
        return 1
    else:
        return 0

def check_west(seat, seats):
    (row, column) = seat
    c = column - 1
    found = False
    while c >= 0 and not found:
        found = is_seat(seats[row][c])
        c -= 1
    if found and seats[row][c + 1] == '#':
        return 1
    else:
        return 0

def check_northwest(seat, seats):
    (row, column) = seat
    r = row - 1
    c = column - 1
    found = False
    while r >= 0 and c >= 0 and not found:
        found = is_seat(seats[r][c])
        r -= 1
        c -= 1
    if found and seats[r + 1][c + 1] == '#':
        return 1
    else:
        return 0

def get_num_occupied(seat, seats):
    (row, column) = seat
    numOccupied = check_north(seat, seats)
    numOccupied += check_northeast(seat, seats)
    numOccupied += check_east(seat, seats)
    numOccupied += check_southeast(seat, seats)
    numOccupied += check_south(seat, seats)
    numOccupied += check_southwest(seat, seats)
    numOccupied += check_west(seat, seats)
    numOccupied += check_northwest(seat, seats)
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
            elif seats[r][c] == '#' and get_num_occupied(seat, seats) >= 5:
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