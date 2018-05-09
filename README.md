### BBoy
(BGirls can use it too)

Request engine wrapped around BoltDB.

3 types of queries ATM:
- Bucket allows to retrieve and update every keys inside a bucket.
- Key allows to retrieve and update a single key inside a bucket.
- MaskBucket allow to retrieve and update sets of keys inside a bucket. map[string]interface{} must be provided as key filter for Update and View. A mode (a uint8) must be provided either to be bboy.Left or bboy.Right value. This defines the behavior during an Update: if Left is selected, Keys passed to the struct **WILL OVERWRITE** Keys inside the targeted Bucket. Else Keys inside the targeted Bucket will always be right (lul).
