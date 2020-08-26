#Basic example of decorator:

def somedecorator(func):
	def wrapper_func(*args):
		for nmr in func(*args):
			print(nmr ** 2)
	return wrapper_func



@somedecorator
def somefunction(*args):
	return args

somefunction(5, 6, 7, 8)
#25
#36
#49
#64