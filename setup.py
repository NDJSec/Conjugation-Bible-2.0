from setuptools import find_packages
from setuptools import setup

setup(
    name="Conjugation Bible 2.0",
    version="2.0",
    description="A program for drilling different conjugations in various romance lanugages",
    author="Taylor Perry",
    packages=find_packages(),
    entry_points={"console_scripts": ["ConjugationBible=PySrc.Romance_Dominus:Romance_Dominus"]}
)