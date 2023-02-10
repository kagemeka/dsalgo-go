import os


def _getext(path: str) -> str:
    return os.path.splitext(path)[-1]


def _get_pkgname(path: str) -> str:
    return os.path.split(path)[-1]


def create_pkg(path: str) -> None:
    assert not _getext(path)
    pkg = _get_pkgname(path)
    assert pkg
    content = f"package {pkg}"
    os.makedirs(path)
    src_file = f"{path}/{pkg}.go"
    test_file = f"{path}/{pkg}_test.go"
    files = [src_file, test_file]
    for p in files:
        with open(p, mode="w") as f:
            f.write(content)


if __name__ == "__main__":
    import argparse

    parser = argparse.ArgumentParser()
    parser.add_argument("pkg_path", type=str)
    args = parser.parse_args()
    create_pkg(args.pkg_path)
