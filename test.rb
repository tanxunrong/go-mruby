
class Sky
end

a = Sky.new

def a.hi()
	puts "hi"
end

puts Sky.singleton_class

b = Sky.new

begin
	a.hi ## singleton method for only one instance
	b.hi
rescue Exception => detail 
	puts detail
end
