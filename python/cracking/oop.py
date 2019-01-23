"""Object-Oriented Stuffs"""
import random
from enum import Enum

# Call Center
class CallCenter:
    def __init__(self, employees, managers, director):
        self._employees = employees
        self._managers = managers
        self._director = director
        self._call_queue = {}

    @property
    def employees(self):
        return self._employees

    @property
    def managers(self):
        return self._managers

    @property
    def director(self):
        return self._director

    def dispatch_call(self, call):
        for employee in self._employees:
            if not employee.busy:
                self._call_queue.update({call, employee})
                return
        for manager in self._managers:
            if not manager.busy:
                self._call_queue.update({call, manager})
                return
        self._director.calls.append(call)


class Employee:
    def __init__(self, name, supervisor):
        self._name = name
        self._busy = False
        self._supervisor = supervisor

    @property
    def busy(self):
        return self._busy

    @property
    def name(self):
        return self._name

    @property
    def supervisor(self):
        return self._supervisor

class Supervisor:
    def __init__(self, employees):
        self._employees = employees

    @property
    def employees(self):
        return self._employees

    def salary(self):
        return 100 * len(self._employees)


class Manager(Employee, Supervisor):
    def __init__(self, name, employees):
        self._calls = []
        super(Manager, self).__init__(name)
        super(Manager, self).__init__(employees)

class Director(Manager, Supervisor):
    def __init__(self, name, managers):
        self._managers = managers
        super(Director, self).__init__(name)

    @property
    def managers(self):
        return self._managers

    @property
    def calls(self):
        return self._calls

# Chat Server
class Server:
    def __init__(self, host):
        self._hostname = host
        self._clients = []
        self._rooms = {}

    @property
    def clients(self):
        return self._clients

    def connect(self, name, host):
        new_client = Client(name, host)
        self._clients.append(new_client)
        conn = Connection(self, new_client)
        return new_client, conn

    @property
    def rooms(self):
        return self._rooms

    def new_room(self, name, client):
        new_room = Room(name)
        new_room.join(client)
        self._rooms.update({name, new_room})
        return new_room

class Room:
    def __init__(self, name):
        self._name = name
        self._messages = []
        self._participants = []

    def join(self, client):
        self._participants.append(client)
        return self

    def leave(self, client):
        self._participants.pop(client)
        return self

    def post(self, message):
        self._messages.append(message)
        return self._messages

class Client:
    def __init__(self, name, host):
        self._name = name
        self._host = host

    @property
    def name(self):
        return self._name

    @property
    def host(self):
        return self._host

class Connection:
    def __init__(self, server, client):
        self._server = server
        self._client = client
        self._queue = []

    def send(self, data):
        self._queue.append(data)
        return data

    def listen(self):
        return self._queue

    def close(self):
        self._queue = []
        return True

# Minesweeper
class Map:
    def __init__(self, size, difficulty):
        self._grid = [[Cell(bool(random.randint(0, 2))) for _ in range(size)] for _ in range(size)]
        self._difficulty = difficulty
        self._is_done = False

    def __repr__(self):
        return "\n".join(["".join(["{}".format(item) for item in row]) for row in self._grid])

    def populate(self):
        return self._grid

    def close(self):
        if self._is_done:
            print("All done")

    def click(self, x, y):
        neighbors = []
        if len(self._grid) <= y or y < 0:
            return False
        if len(self._grid[y]) <= x or x < 0:
            return False
        if y > 0:
            if x > 0:
                neighbors.append(self._grid[y][x-1])
                neighbors.append(self._grid[y-1][x-1])
            if x < len(self._grid[y]) - 1:
                neighbors.append(self._grid[y][x+1])
                neighbors.append(self._grid[y-1][x+1])
            neighbors.append(self._grid[y-1][x])
        if y < len(self._grid) - 1:
            neighbors.append(self._grid[y+1][x])
            if x > 0:
                neighbors.append(self._grid[y+1][x-1])
            if x < len(self._grid[y]) - 1:
                neighbors.append(self._grid[y+1][x+1])
        return self._grid[y][x].click(neighbors)

class Marks(Enum):
    NONE = "none"
    FLAG = "flag"
    BOMB = "bomb"

    def __repr__(self):
        return "{}".format(self.name)

    def __str__(self):
        return "{}".format(self.name)

class Cell:
    marks = Enum("none", "flag")

    def __init__(self, is_bomb):
        self._is_bomb = is_bomb
        self._cover = True
        self._number = 0
        self._mark = Marks.NONE

    def __repr__(self):
        if self._cover:
            return "{:^6}".format(self._mark)
        if self._is_bomb:
            return "{:^6}".format(Marks.BOMB)
        return "{:^6}".format(self._number)

    def __str__(self):
        return self.__repr__()

    def click(self, neighbors):
        if self._cover:
            self._cover = False
        if self._is_bomb:
            print("Asplode")
            self._mark = Marks.BOMB
            return self._mark
        bombs_nearby = list(filter(lambda x: x.check(), neighbors))
        self._number = len(bombs_nearby)
        return self._number

    def check(self):
        return self._is_bomb

    @property
    def number(self):
        return self._number

    @property
    def mark(self):
        return self._mark
