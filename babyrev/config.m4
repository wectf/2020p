PHP_ARG_ENABLE(babyrev, whether to enable babyrev support, 
[ --enable-babyrev   Enable babyrev support])

if test "$PHP_BABYREV" = "yes"; then
   AC_DEFINE(HAVE_BABYREV, 1, [Whether you have babyrev])
   PHP_NEW_EXTENSION(babyrev, babyrev.c, $ext_shared)
fi
