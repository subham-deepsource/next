import utils, mock

def simplifable_if_statement(arg1, arg2):
    # Can be reduced to bool(arg1 and not arg2)
    if arg1 and not arg2: # [simplifiable-if-statement]
        return True
    else:
        return False

def test_simplifiable_if_expreseeion(arg):
    #  Can be replaced by bool(arg)
    return True if arg else False # [simplifiable-if-expression]
