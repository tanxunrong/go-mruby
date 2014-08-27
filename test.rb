
class Sky
end

a = Sky.new

def a.hi()
	puts "hi"
end

puts Sky.singleton_class

b = Sky.new

c = [1,2]
d = [1,2]
puts c == d
puts c.eql? d
puts c.equal? d

begin
	a.hi ## singleton method for only one instance
	b.hi
rescue Exception => detail 
	puts detail
end
