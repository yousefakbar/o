PACKAGE		= o
PACKAGE_VERSION	= dev
PACKAGE_DIR	= $(shell dirname "$0")

SRCFILES	= cmd/o/main.go
BINFILES	= $(PACKAGE)

GO		= go
INSTALL		= install -c
UNINSTALL	= rm -f
RM		= rm

bindir		= /usr/local/bin

######################################################################

build: $(SRCFILES)
	@echo Building: $(SRCFILES)
	@$(GO) build -o $(PACKAGE) cmd/o/main.go

install:
	@for f in $(BINFILES); do \
		echo Installing: $$f; \
		$(INSTALL) $$f $(DESTDIR)$(bindir)/`basename $$f` || exit 1; \
	done

uninstall:
	@for f in $(BINFILES); do \
		echo Uninstalling: $$f; \
		$(UNINSTALL) $(DESTDIR)$(bindir)/`basename $$f`; \
	done

clean:
	@for f in $(BINFILES); do\
		echo Removing: $$f; \
		$(RM) `basename $$f`; \
	done

fullclean: uninstall clean
