if ARGV.count != 1
  exit
end

File.readlines(ARGV[0]).each do |line|
  count = 0
  num = line.to_i

  while num != 0 do
    if (num & 1) == 1
      count += 1
    end

    num = num >> 1
  end

  puts count
end
