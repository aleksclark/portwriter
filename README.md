# portwriter
port writing wrapper for sending bytes to different places.

This is intented to provide a common interface to byte-receiving peripherals such as GPIOs using bit-banging, I2C, or SPI.

Currently only I2C is implemented, and that leverages the sharedi2c package for https://github.com/kidoman/embd
