import time
import unittest
from typing import Optional

from pytnon.ctci.c2.LinkedList import LinkedList


# DOG = "dog"
# CAT = "cat"


class Animal:
    def __init__(self, id_val):
        self.created_at = time.time()
        self.id = id_val


class Dog(Animal):
    def __init__(self, id):
        Animal.__init__(self, id)


class Cat(Animal):
    def __init__(self, id):
        Animal.__init__(self, id)


class AnimalShelter:

    def __init__(self):
        self.dogs = LinkedList()
        self.cats = LinkedList()

    def enqueue(self, animal: Animal):
        if isinstance(animal, Dog):
            self.dogs.add(animal)
        else:
            self.cats.add(animal)

    def dequeue_any(self) -> Optional[Animal]:
        if not self.dogs.head:
            return self.__get_first(self.cats)

        if not self.cats.head:
            return self.__get_first(self.dogs)

        # pick the oldest created
        if self.cats.head.value.created_at < self.dogs.head.value.created_at:
            return self.__get_first(self.cats)

        return self.__get_first(self.dogs)

    def dequeue_dog(self) -> Optional[Dog]:
        return self.__get_first(self.dogs)

    def dequeue_cat(self) -> Optional[Cat]:
        return self.__get_first(self.cats)

    def __get_first(self, list: LinkedList) -> Optional[Animal]:
        if not list.head:
            return None

        first = list.head
        list.head = list.head.next
        return first.value


class Test(unittest.TestCase):
    def test_animal_shelter(self):
        shelter = AnimalShelter()

        shelter.enqueue(Cat(1))
        shelter.enqueue(Cat(2))
        shelter.enqueue(Cat(3))

        self.assertAnimal(Cat(1), shelter.dequeue_any())
        self.assertAnimal(Cat(2), shelter.dequeue_cat())
        self.assertAnimal(None, shelter.dequeue_dog())

        shelter.enqueue(Dog(1))
        shelter.enqueue(Dog(2))
        shelter.enqueue(Dog(3))

        self.assertAnimal(Dog(1), shelter.dequeue_dog())
        self.assertAnimal(Cat(3), shelter.dequeue_any())
        self.assertAnimal(Dog(2), shelter.dequeue_any())
        self.assertAnimal(Dog(3), shelter.dequeue_any())
        self.assertAnimal(None, shelter.dequeue_any())
        self.assertAnimal(None, shelter.dequeue_any())
        self.assertAnimal(None, shelter.dequeue_cat())
        self.assertAnimal(None, shelter.dequeue_dog())

    def assertAnimal(self, want: Optional[Animal], actual: Optional[Animal]):
        if actual is None or want is None:
            self.assertEqual(want, actual)
            return

        self.assertEqual(want.id, actual.id)
        # print(type(want), type(actual))
        self.assertEqual(True, type(want) is type(actual))
