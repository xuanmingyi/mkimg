# 配置文件

```yaml
output: hello.img

copy:
    - file: boot.bin
      offset: 0
```

* output 软驱镜像1440kb(1.44mb)
* boot 是启动分区, 0-512字节