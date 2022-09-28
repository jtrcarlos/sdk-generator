<?php

namespace Tests;

class SwiftServer56Test extends Base
{
    protected string $language = 'swift-server';
    protected string $class = 'Appwrite\SDK\Language\Swift';
    protected array $build = [
        'mkdir -p tests/sdks/swift-server/Tests/AppwriteTests',
        'cp tests/languages/swift-server/Tests.swift tests/sdks/swift-server/Tests/AppwriteTests/Tests.swift',
    ];
    protected string $command =
        'docker run --rm -v $(pwd):/app -w /app/tests/sdks/swift-server swift:5.6.1-focal swift test';

    protected array $expectedOutput = [
        ...Base::FOO_RESPONSES,
        ...Base::BAR_RESPONSES,
        ...Base::GENERAL_RESPONSES,
        ...Base::LARGE_FILE_RESPONSES,
        ...Base::EXTENDED_GENERAL_RESPONSES,
        ...Base::EXCEPTION_RESPONSES,
    ];
}
