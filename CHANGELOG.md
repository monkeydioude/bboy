### 3.0.0 - 06-05-2019

- Changed constructor + added helpers func

### 2.0.0 - 03-05-2019

- Moved to github.com/coreos/bbolt

### 1.1.2 - 14-06-2018

- Add error handling on key/view

### 1.1.1 - 10-06-2018

- Modified logging message in View

### 1.1.0 - 09-06-2018

- Use coreos/bbolt package now, which is a fork of bolt
- Add and use SafeCopy function, protecting from unsafe []byte pointer manipulation from bolt

### 1.0.2 - 20-05-2018

- Better logging on db couldnt be opened

### 1.0.1 - 09-05-2018

- MODIFIED: Mode inside MaskBucket is now uint8 (int8 before)
- ADDED: CHANGELOG & README

### 1.0.0 - 09-05-2018

- Added requesting engine for BoltDB
- Bucket Query
- Key Query
- MaskBucket Query
