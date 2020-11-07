module hello

go 1.15

replace kaka.com/maxConsecutiveOnes => ../maxConsecutiveOnes

replace kaka.com/greetings => ../greetings

require (
	kaka.com/greetings v0.0.0-00010101000000-000000000000
	kaka.com/maxConsecutiveOnes v0.0.0-00010101000000-000000000000
)
