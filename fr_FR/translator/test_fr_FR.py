import pytest

from translator import fr_FR


def test_en_to_fr_hi():
    result = fr_FR.en_to_fr("hi")
    assert len(result) > 0


def test_en_to_fr_bye():
    result = fr_FR.en_to_fr("bye")
    assert len(result) > 0
