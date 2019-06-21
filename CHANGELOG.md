|Date      |Issue |Description                                                                                              |
|----------|------|---------------------------------------------------------------------------------------------------------|
|2019/06/21|      |Release 0.4.3                                                                                            |
|2019/06/21|      |Support `go mod`                                                                                         |
|2019/06/12|      |Release 0.4.2                                                                                            |
|2019/04/20|35    |Support email SANs in certificate names                                                                  |
|2019/05/29|43    |Log privileged cert caches only when the cache is changed                                                |
|2019/05/23|      |Release 0.4.0                                                                                            |
|2019/05/23|40    |Fix SecurityAlwaysOverwriteCache handling for the Puppet provider                                        |
|2019/05/22|38    |Should a user cert AND a privileged cert exist, check user cert first                                    |
|2019/05/22|36    |When SecurityAlwaysOverwriteCache is set, only write when the certs actually change                      |
|2019/01/17|      |Release 0.3.0                                                                                            |
|2019/01/03|27    |Validate privileged certificates using their expected name instead of the claimed caller                 |
|2019/01/02|28    |Allow callerid schemes other than `choria=`                                                              |
|2019/01/02|29    |When checking cert validations check privileged ones first                                               |
|2018/11/23|      |Release 0.2.1                                                                                            |
|2018/11/23|22    |Improve compatibility with Ruby Choria by supporting slashes in regex                                    |
|2018/11/15|      |Release 0.2.0                                                                                            |
|2018/11/11|18    |Allow caching to always overwrite cached certificates to deal with short lived certificates              |
|2018/11/11|16    |Validate certificates before caching to ensure future policy changes are re-evaluated                    |
|2018/11/06|13    |Support intermediate certificates during validation                                                      |
|2018/06/15|      |Release 0.1.0                                                                                            |
|2018/06/15|8     |Do not confuse the concepts of a `certname` and `identity`                                               |
|2018/05/30|      |Release 0.0.2                                                                                            |
|2018/05/30|5     |Handle errors when passing options                                                                       |
|2018/05/23|      |Release 0.0.1                                                                                            |
|2018/05/23|      |Extract from the `go-choria` project                                                                     |