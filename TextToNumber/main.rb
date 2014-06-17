def numeric(text)
  case text.downcase
  when 'zero'
    0
  when 'one'
    1
  when 'two'
    2
  when 'three'
    3
  when 'four'
    4
  when 'five'
    5
  when 'six'
    6
  when 'seven'
    7
  when 'eight'
    8
  when 'nine'
    9
  when 'ten'
    10
  when 'eleven'
    11
  when 'twelve'
    12
  when 'thirteen'
    13
  when 'fourteen'
    14
  when 'fifteen'
    15
  when 'sixteen'
    16
  when 'seventeen'
    17
  when 'eighteen'
    18
  when 'nineteen'
    19
  when 'twenty'
    20
  when 'thirty'
    30
  when 'forty'
    40
  when 'fifty'
    50
  when 'sixty'
    60
  when 'seventy'
    70
  when 'eighty'
    80
  when 'ninety'
    90
  when 'hundred'
    100
  when 'thousand'
    1000
  when 'million'
    1000000
  else
    nil
  end
end

def words_to_num(words)
  if words.count == 0
    return 0
  end

  largest = -1
  largest_index = -1

  words.each_with_index do |word, index|
    num = numeric(word)

    if num > largest
      largest = num
      largest_index = index
    end
  end

  result = largest

  multiplier_words = words[0...largest_index]
  if multiplier_words.count > 0
    result *= words_to_num(multiplier_words)
  end

  rest_of_words = words[(largest_index+1)..-1]

  return result + words_to_num(rest_of_words)
end

def parse_number(text)
  words = text.split(' ')

  if words[0] == 'negative'
    return -1 * words_to_num(words[1..-1])
  else
    return words_to_num(words)
  end
end

if ARGV.count != 1
  exit
end

File.readlines(ARGV[0]).each do |line|
  puts parse_number(line.chomp)
end
