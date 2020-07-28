

def en_to_fr(word: str) -> str:
    """
    Translator english to french
    :param word:
    :return:
    """

    if word.lower() in ["hi", "hello"]:
        return "bonjour"
    elif word.lower() == "bye":
        return "adieu"
    else:
        return ""