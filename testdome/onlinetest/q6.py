# from collections import namedtuple


# def merge(*records):
#     """
#     :param records: (varargs list of namedtuple) The patient details.
#     :returns: (namedtuple) named Patient, containing details from all records, in entry order.
#     """
#     return None


# PersonalDetails = namedtuple('PersonalDetails', ['date_of_birth'])
# personal_details = PersonalDetails(date_of_birth='06-04-1972')

# Complexion = namedtuple('Complexion', ['eye_color', 'hair_color'])
# complexion = Complexion(eye_color='Blue', hair_color='Black')

# print(merge(personal_details, complexion))

from collections import namedtuple

# Patient = namedtuple('Patient', ['eye_color', 'hair_color', 'date_of_birth'])


def merge(*records):
    """
    :param records: (varargs list of namedtuple) The patient details.
    :returns: (namedtuple) named Patient, containing details from all records, in entry order.
    """
    fields = [field for record in records for field in record._fields]

    Patient = namedtuple('Patient', fields)

    merged = {k: d[k] for d in [r._asdict() for r in records] for k in d}
    return Patient(**merged)


PersonalDetails = namedtuple('PersonalDetails', ['date_of_birth'])
personal_details = PersonalDetails(date_of_birth='06-04-1972')


Complexion = namedtuple('Complexion', ['eye_color', 'hair_color'])
complexion = Complexion(eye_color='Blue', hair_color='Black')

print(merge(personal_details, complexion))
print(merge(complexion, personal_details))
