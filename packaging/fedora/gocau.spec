%global debug_package %{nil}

Name:           gocau
Version:        1.0.0
Release:        1%{?dist}
Summary:        Terminal calculator CLI — avg, weighted avg, fitness metric

License:        MIT
URL:            https://github.com/xieguaiwu/gocau
Source0:        %{url}/archive/v%{version}.tar.gz#/%{name}-%{version}.tar.gz

BuildRequires:  golang >= 1.24

%description
gocau is a minimal terminal calculator CLI with three subcommands:
arithmetic average, weighted average, and a linguistic Fitness Metric.
Zero third-party dependencies (Go standard library only).

%prep
%setup -q -n gocau-%{version}

%build
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

go build -trimpath -ldflags="-s -w" -o gocau .

%install
rm -rf %{buildroot}
install -Dm755 gocau %{buildroot}%{_bindir}/gocau
install -Dm644 LICENSE %{buildroot}%{_defaultlicensedir}/%{name}/LICENSE
install -Dm644 README.md %{buildroot}%{_defaultdocdir}/%{name}/README.md

%files
%license LICENSE
%doc README.md
%{_bindir}/gocau

%changelog
* Tue Aug 18 2026 xgw <xieguaiwu@163.com> - 1.0.0-1
- Initial package: terminal calculator CLI (v1.0.0)
