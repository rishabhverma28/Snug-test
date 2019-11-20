#!/usr/bin/env python3

import argparse

YEAR_MIN = 1901
YEAR_MAX = 2999

month_days_leap = [31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
month_days = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]

cumm_days = [0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334]
cumm_days_leap = [0, 31, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335]

'''
Function to check if the year is a leap year.
'''


def check_leap_year(year):
    if(year % 4 == 0 and year % 100 != 0) or (year % 400 == 0):
        return True
    return False


'''
Class that's used in the App to throw exceptions
and print the error.
'''


class AppException(argparse.ArgumentTypeError):
    def __init__(self, data):
        self.data = data

    def __str__(self):
        return repr(self.data)


'''
Class to initialize the date.

Instance Variables:
day, month, year, is_leap_year, since_origin

Methods:
form_valid_date: carry out the operations to fill the 
                instance variables
check_valid_month_year: Check if the month and Year are valid.
check_valid_day: Check if the day is valid.
num_days_since_origin: Calculate the day since 01/01/0001
'''


class NewDate():
    def __init__(self, date):
        self.form_valid_date(date)
        self.since_origin = self.num_days_since_origin()

    def form_valid_date(self, date):
        date_string = date.split("/")
        # Check if day, month, year fields are of the correct length
        if(len(date_string) != 3 or len(date_string[0]) != 2 or len(date_string[1]) != 2 or len(date_string[2]) != 4):
            raise AppException(
                "Please Make sure that the date is passed in the correct format. [DD/MM/YYYY]")
        date_list = []
        # Convert all three fields to integers
        for i in date_string:
            b = int(i)
            date_list.append(b)
        self.day, self.month, self.year = date_list
        self.is_leap_year = check_leap_year(self.year)
        # Check all fields are within the correct range
        valid_date = self.check_valid_month_year()
        if not valid_date:
            raise AppException(
                "Invalid Date. Please Make sure that the date passed is correct")

    def check_valid_day(self):
        # Take into account if the year is a leap year
        if self.is_leap_year:
            if self.day not in range(1, month_days_leap[self.month-1]):
                return False
        else:
            if self.day not in range(1, month_days[self.month-1]):
                return False
        return True

    def check_valid_month_year(self):
        # Check if the month and year are in the correct range
        if self.month not in range(1, 13):
            return False
        if self.year not in range(YEAR_MIN, YEAR_MAX+1):
            return False
        # Check if the day is valid
        valid_day = self.check_valid_day()
        if not valid_day:
            return False
        return True

    def num_days_since_origin(self):
        # Calculate the cummulative days in the year ranging up
        # to the current day in the year.
        # Dependent upon the year. If year is a leap year,
        # then an extra day is added
        if self.is_leap_year:
            present_day = self.day + cumm_days_leap[self.month - 1]
        else:
            present_day = self.day + cumm_days[self.month - 1]
        # Calculate the number of leap years before the present year
        numLeapYears = (self.year-1)//4 - \
            (self.year-1)//100 + (self.year-1)//400
        # return the number of days which is the sum of
        # number of days of the present year the number of leap years
        # before the present year as these will be the extra days
        # needed to be added to the (previous years)* 365
        return (self.year-1)*365 + numLeapYears + present_day - 1


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description=f"Calculate the number of full days \
                                     elapsed between the experiment's start date and end date. \
                                     Enter valid dates between 01/01/{YEAR_MIN} and 31/12/{YEAR_MAX}.")
    parser.add_argument("date1", help="Experiment Start Date")
    parser.add_argument("hyphen", help="The hyphen", default="-")
    parser.add_argument("date2", help="Experiment End Date")
    args = parser.parse_args()
    # Initialize and Validate the input dates
    dd1 = NewDate(args.date1)
    dd2 = NewDate(args.date2)
    # Print out the result
    print(max(abs(dd2.since_origin - dd1.since_origin) - 1, 0))
